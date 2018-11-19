package main

import (
	"net"
	"fmt"
	"io"
)

func main() {
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buf := make([]byte, 5)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		fmt.Println(string(buf))

		n, err := conn.Write(buf)
		if err != nil {
			panic(err)
		}
		if n != len(buf) {
			panic("n != len(bs)")
		}
	}
}
