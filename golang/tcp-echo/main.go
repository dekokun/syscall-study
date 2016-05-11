package main

import (
	"fmt"
	"os"
	"net"
	"io/ioutil"
	"time"
)

func main() {
	usage := fmt.Sprintf("Usage: $s client|server host:port\n", os.Args[0])
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, usage)
		os.Exit(1)
	}
	service := os.Args[2]

	if os.Args[1] == "client" {
		client(service)
	} else if os.Args[1] == "server" {
		server(service)
	} else {
		fmt.Fprintf(os.Stderr, usage)
		os.Exit(1)
	}

}



func dieIfError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
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
	fmt.Println(string(result))
	os.Exit(0)
}

func server(service string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	dieIfError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	dieIfError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now().String()
		fmt.Println("Access come !")
		conn.Write([]byte(daytime))
		conn.Close()
	}

}
