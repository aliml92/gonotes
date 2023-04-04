package main


import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {
	urlExample := "postgres://demouser:demopassword@localhost:5432/conduitdb?sslmode=disable"
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	pgCommand, err := conn.Exec(context.Background(), "create table if not exists widgets (id serial primary key, name text, weight integer)")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create table: %v\n", err)
		os.Exit(1)
	}
	i := pgCommand.RowsAffected()
	fmt.Println(i)
	pgCommand, err = conn.Exec(context.Background(), "insert into widgets (name, weight) values ($1, $2)", "widget name", 10)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to insert into table: %v\n", err)
		os.Exit(1)
	}
	i = pgCommand.RowsAffected()
	fmt.Println(i)
	
	var name string
	var weight int64
	err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 2).Scan(&name, &weight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name, weight)
}