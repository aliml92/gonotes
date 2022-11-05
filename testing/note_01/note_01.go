// link: https://blog.logrocket.com/advanced-unit-testing-patterns-go/
package main

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

// index accepts http.ResponseWriter interface and
// *http.Request as parameters and writes 200 to Header
func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}


// getJoke uses external dependency
func getJoke(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jokeId := u.Query().Get("id")
	if jokeId == "" {
		http.Error(w, "Joke ID cannot be empty", http.StatusBadRequest)
		return
	}

	endpoint := "https://icanhazdadjoke.com/j/" + jokeId

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("Accept", "text/plain")

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != http.StatusOK {
		http.Error(w, string(b), resp.StatusCode)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write(b) 
}


func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/joke", getJoke)
	http.ListenAndServe(":1212", mux)
}