package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const createTable = `CREATE TABLE IF NOT EXISTS big_numbers (id serial PRIMARY KEY, value bigint)`

const insert = `INSERT INTO big_numbers (value) VALUES ($1) RETURNING id`

const update = `
UPDATE big_numbers 
	SET value = COALESCE(NULLIF($2, 0), value)
WHERE id = $1 
RETURNING id, value 
`

type BigNumber struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}



func main() {

	dsn := "host=localhost port=5435 user=demouser password=demopassword dbname=demodb sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ctx := context.Background()

	// create table
	_, err = db.Exec(createTable)
	if err != nil {
		panic(err)
	}

	// insert initial data
	var id int64
	err = db.QueryRowContext(ctx, insert, 10).Scan(&id)
	if err != nil {
		panic(err)
	}


	// update data (works)
	newValue := int64(2147483647)
	u, err := UpdateUser(ctx, db, id, newValue)  // 2147483647 is the max value for int32
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Updated: %v\n", u)

	// update data (fails)
	newValue = int64(2147483648)
	_, err = UpdateUser(ctx, db, id, newValue)  
	if err != nil {
		fmt.Printf("Failed to update: %v\n", err)
	}
}

func UpdateUser(ctx context.Context, db *sql.DB, id int64, value int64) (BigNumber, error) {
	row := db.QueryRowContext(ctx, update, id, value)
	var b BigNumber
	err := row.Scan(&b.ID, &b.Value)
	return b, err
}