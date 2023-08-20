package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int) // ch is unbuffered channel
	go func(){
		fmt.Println("child goroutine is running")
		time.Sleep(5 * time.Second) // simulates a long running function
		fmt.Println("sending signal through channel")
		ch <- 1
		time.Sleep(5 * time.Second) // simulates a long running function
		ch <- 2
	}()
	fmt.Println("waiting for the signal from child goroutine")
	sig, ok := <-ch
	if ok {
		fmt.Println("channel is open")
	} else {
		fmt.Println("channel is closed")
	}
	fmt.Printf("signal received: %d\n", sig)
	s, ok := <-ch
	if ok {
		fmt.Println("channel is still open")
	} else {
		fmt.Println("channel is closed")
	}
	fmt.Printf("signal received: %d\n", s)
} 