package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"net/http"
	"time"

	"github.com/k0kubun/pp/v3"
)

func main() {



	go func(){
		mux := http.NewServeMux()
		mux.HandleFunc("/fast", func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(1 * time.Second)
			fmt.Fprintf(w, "client request url: %s\n", r.URL)
		})
		mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(6 * time.Second)
			fmt.Fprintf(w, "client request url: %s\n", r.URL)
		})
		mux.HandleFunc("/super-slow", func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(20 * time.Second)
			fmt.Fprintf(w, "client request url: %s\n", r.URL)
		})
		http.ListenAndServe(":8010", mux)
	}()

	// this does not return error
	if err := runClient(nil, "/fast"); err != nil {
		pp.Printf("fast endpoint: %v\n", err)
	}

	// this client req shows deadline exceeded error
	if err := runClient(nil, "/slow"); err != nil {
		pp.Printf("slow endpoint: %v\n", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up a channel to listen for SIGINT (Ctrl+C)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	errChan := make(chan error, 1)
	go func(ctx context.Context) {
		if err := runClient(ctx, "/super-slow"); err != nil {
			errChan <- fmt.Errorf("super slow endpoint,\n couldn't wait: %v\n", err)
		}
	}(ctx)

	select {
	case <-sigChan:
		// If we receive a SIGINT, cancel the context
		cancel()
		fmt.Println("Received SIGINT, canceling request...")
		err := <- errChan
		pp.Printf("err after canceling request: %v\n", err)
	case err := <-errChan:
		// If we receive a err first, print it
		pp.Printf("err after exceeding deadline: %v\n", err)
	}	


}


func runClient(ctx context.Context, path string) error {
	
	url := "http://localhost:8010/" + path
	
	client := http.DefaultClient

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()

	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Convert the body to string for easier display
	bodyString := string(bodyBytes)
	
	fmt.Printf("Response status: %s\n", resp.Status)
	fmt.Printf("Reponse body: %s\n", bodyString)
	return nil
}