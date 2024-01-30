package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main(){

	// Dial function connects to a server
	// and used by client to make request
	go func(){
		conn, err := net.Dial("tcp", ":8010")
		if err != nil {
			log.Fatalf("Failed to dial: %v\n", err)
		}
		defer conn.Close()

		// sending request to client
		fmt.Fprintf(conn, "hello server, give me an apple\r\n")
		
		// reading response from server
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read: %v\n", err)
		}
		log.Printf("Server response: %v\n", response)
	}()

	// Listen function creates servers
	ln, err := net.Listen("tcp", ":8010")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("Failed to accept conn: %v\n", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// reading client request
	request, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read: %v\n", err)
	}
	log.Printf("client request: %v\n", request)

	// sending response to client
	fmt.Fprintf(conn, "hello client, here is your 🍎\r\n")
}