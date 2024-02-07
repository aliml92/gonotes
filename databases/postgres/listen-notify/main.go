package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	// create db connection pool
	dbpool, err := pgxpool.New(context.Background(), "postgres://lbdbuser:lndbpassword@localhost:5434/lndb?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	defer dbpool.Close()

	


}