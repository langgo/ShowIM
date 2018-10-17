package main

import (
	"net"
	"fmt"
)

func main() {
	bs := []byte("hello")

	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		panic(err)
	}

	fmt.Println("1")
	_, err = conn.Write(bs)
	if err != nil {
		panic(err)
	}

	fmt.Println("2")
	buf := make([]byte, 5)
	_, err = conn.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Println("return=>", string(buf))
	if err := conn.Close(); err != nil {
		panic(err)
	}
}
