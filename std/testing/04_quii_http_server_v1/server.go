package main

import (
	"fmt"
	"net/http"
)

// PlayerServer is an only guy who is gonna be resposible for fates of all future requests
// to my insanely dumb server, because it does not care about what a client *Request
// wants from my server, instead it just sends "20" back to all clients
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}  
