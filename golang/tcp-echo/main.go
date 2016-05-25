package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	usage := fmt.Sprintf("Usage: $s client|server host:port\n", os.Args[0])
	if len(os.Args) != 3 {
		log.Fatal(usage)
	}
	service := os.Args[2]

	if os.Args[1] == "client" {
		client(service)
	} else if os.Args[1] == "server" {
		server(service)
	} else {
		log.Fatal(usage)
	}
}

func dieIfError(err error) {
	if err != nil {
		log.Fatal("Fatal error: %s", err.Error())
	}
}

func client(service string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	dieIfError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	dieIfError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	dieIfError(err)
	result, err := ioutil.ReadAll(conn)
	dieIfError(err)
	log.Println(string(result))
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
		now := time.Now().String()
		log.Println(strconv.Itoa(count) + "Access come !")
		conn.Write([]byte(now))
		conn.Close()
	}

}
