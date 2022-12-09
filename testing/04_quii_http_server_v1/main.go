// source: https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server
package main

import (
	"log"
	"net/http"
)


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