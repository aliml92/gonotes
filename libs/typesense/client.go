package main

import (
	"fmt"

	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
	"github.com/typesense/typesense-go/typesense/api/pointer"
)


func main() {
	// create client
	client := typesense.NewClient(
	    typesense.WithServer("http://localhost:8110"),
	    typesense.WithAPIKey("xyz"))

	// create schema	
	schema := &api.CollectionSchema{
		Name: "companies",
		Fields: []api.Field{
			{
				Name: "company_name",
				Type: "string",
			},
			{
				Name: "num_employees",
				Type: "int32",
			},
			{
				Name:  "country",
				Type:  "string",
				Facet: pointer.True(),
			},
		},
		DefaultSortingField: pointer.String("num_employees"),
	}

	client.Collections().Create(schema)
	
	// create document
	document := struct {
		ID           string `json:"id"`
		CompanyName  string `json:"company_name"`
		NumEmployees int    `json:"num_employees"`
		Country      string `json:"country"`
	}{
		ID:           "123",
		CompanyName:  "Stark Industries",
		NumEmployees: 5215,
		Country:      "USA",
	}

	client.Collection("companies").Documents().Create(document)

	// update document
	d := struct {
		CompanyName  string `json:"company_name"`
		NumEmployees int    `json:"num_employees"`
	}{
		CompanyName:  "Stark Industries",
		NumEmployees: 5500,
	}

	res, err := client.Collection("companies").Document("123").Update(d)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("res: %v\n", res)
}