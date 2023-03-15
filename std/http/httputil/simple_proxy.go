package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)


func main(){

	r := gin.Default()

	r.GET("/api/v1/flights", fetchFlights)
	
	r.Run(":8080")
}


func fetchFlights(c *gin.Context) {
	remote, err := url.Parse("http://myremotedomain.com")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(r *http.Request) {
		r.URL.Scheme = remote.Scheme
		r.URL.Host = remote.Host
		r.URL.Path = "/flights"
		r.Header = c.Request.Header
		r.Host = remote.Host
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}