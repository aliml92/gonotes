package main

import (
	"fmt"
	"net"
)

func main() {
	go runClint()
	runServer()
}

func runClint() {
	conn, err := net.Dial("tcp", "localhost:8010")
	if err != nil {
		panic(err)
	}
	// send message to server
	message := []byte("hello server")
	_, err = conn.Write(message)
	if err != nil {
		fmt.Printf("failed to write: %v\n", err)
	}

	// read server response
	buf := make([]byte, 1024)
	len, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("failed to read: %v\n", err)
	}
	fmt.Printf("server response: %s\n", string(buf[:len]))
} 

func runServer() {
	addr := "localhost:8010"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	host, port, err := net.SplitHostPort(l.Addr().String())
	if err != nil {
		panic(err)
	}
	fmt.Printf("listening on host: %s, port: %s\n", host, port)

	for {
		// listen for incoming connection
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

		go func(conn net.Conn){
			
			// read client request
			buf := make([]byte, 1024)
			len, err := conn.Read(buf)
			if err != nil {
				fmt.Printf("failed to read: %#v\n", err)
				return
			}
			fmt.Printf("client request: %s\n", string(buf[:len]))

			// send response to client
			conn.Write([]byte("hello client.\n"))
			conn.Close()
		}(conn)
	}
}