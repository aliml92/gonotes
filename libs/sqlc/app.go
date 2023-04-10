package main

import (
	"context"
	"database/sql"
	"log"
	"reflect"

	"github.com/aliml92/gonotes/sqlc/db"

	_ "github.com/lib/pq"
)


func run() error {
	ctx := context.Background()

	// Create a new database connection pool
	postgresdb, err := sql.Open("postgres", "postgres://devuser:devpass@localhost:5432/testdb?sslmode=disable")
	if err != nil {
		return err
	}
    
	queries := db.New(postgresdb)

	// Create a new author
	insertedAuthor, err := queries.CreateAuthor(ctx, db.CreateAuthorParams{
		Bio: "I'm a software developer",
		BirthYear: 1990,
	})
	if err != nil {
		return err
	}
	log.Println("Inserted author:", insertedAuthor)

	// Get the author we just created
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		return err
    }

	log.Println("Fetched author:", fetchedAuthor)
	log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))
	return nil

}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}