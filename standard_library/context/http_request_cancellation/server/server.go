package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func main(){
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Second)
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))	
	})
	log.Fatal(http.ListenAndServe(":8030", nil))
}