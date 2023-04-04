package main

import (
	"context"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)


func main(){
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	log.Println(elasticsearch.Version)
	log.Println(es.Info())

	res, err := es.Index(
		"test",
		strings.NewReader(`{"title" : "Test" }`),
		es.Index.WithDocumentID("1"),
		es.Index.WithRefresh("true"),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	log.Println(res)

	req := esapi.IndexRequest{
		Index:      "test",
		Body: 	 strings.NewReader(`{"title" : "Test" }`),
		DocumentID: "1",
		Refresh:    "true",
	}

	res, err = req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	log.Println(res)
}