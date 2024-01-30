package main

import (
	"context"
	"io"
	"log"
	"net/http"

)


func main(){

	res, err :=  GetResourse(context.Background())
	if err != nil {
		log.Panicln(err)
	}
	log.Println(res)
}

func GetResourse(ctx context.Context) (string, error) {
	client := http.DefaultClient
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8030/foo", nil)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	req = req.WithContext(ctx)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}