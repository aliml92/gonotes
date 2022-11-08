// source: https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server
package main

import (
	"fmt"
	"log"
	"net/http"
)

// PlayerServer is an only guy who is gonna be resposible for fates of all future requests
// to my insanely dumb server, because it does not care about what a client *Request 
// wants from my server, instead it just sends "20" back to all clients
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}  

func main() {

	// here I am creating a variable with a base type of func(ResponseWriter, *Request)
	// If I carefully see what I am doing here, I am converting PlayerServer function
	// to HandlerFunc type, which has the same signature
	// At this moment, I can convince myself that this handler is the handler that 
	// singlehandedly handles all future requests to my server
	// Because whenever a request comes in, HandlerFunc(f) is invoked 
	handler := http.HandlerFunc(PlayerServer)


	// Here I can pass handler to ListenAndServe
	// Now, I am passing handler with a type of http.HandleFunc
	// which in turn has a base type of func(ResponseWriter, *Request)
	// and also implements ServeHTTP method.
	log.Fatal(http.ListenAndServe(":5000", handler))
}