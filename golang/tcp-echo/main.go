package main

import (
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)
import (
	"runtime/pprof"
	"runtime"
	"flag"
)

func main() {
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to file")
	commandtype := flag.String("type", "", "server or client")
	service := flag.String("service", "", "like :8080")
	times := flag.Int("times", 1, "client exec times")
	flag.Parse()
	runtime.SetBlockProfileRate(1)
	log.Println(*cpuprofile)
	if *cpuprofile != "" {
		log.Println("cpuprofiling")
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	if *commandtype == "client" {
		client(*service, *times)
	} else if *commandtype == "server" {
		server(*service)
	} else {
		log.Fatal("not exists type: " + *commandtype)
		log.Fatal(flag.Usage)
	}
}

func dieIfError(err error) {
	if err != nil {
		log.Fatal("Fatal error: %s", err.Error())
	}
}

func client(service string, times int) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	dieIfError(err)
	for i := 0; i < times; i++ {
		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		dieIfError(err)
		_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
		dieIfError(err)
		result, err := ioutil.ReadAll(conn)
		dieIfError(err)
		log.Println(string(result))
	}
}

func server(service string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	dieIfError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	dieIfError(err)
	count := 0
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		count += 1
		responce(conn, count)
		if count == 10 {
			return
		}
	}

}

func responce(conn net.Conn, count int) {
	defer conn.Close()
	now := time.Now().String()
	log.Println(strconv.Itoa(count) + " Access come !")
	conn.Write([]byte(now))
}
