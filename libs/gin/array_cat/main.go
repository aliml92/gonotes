package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

const createTable = `CREATE TABLE IF NOT EXISTS users (id serial PRIMARY KEY, hobbies text[])`

const insert = `INSERT INTO users (hobbies) VALUES ($1) RETURNING id`

const update = `
UPDATE users 
	SET hobbies = array_cat(hobbies, $2)
WHERE id = $1 
RETURNING id, hobbies 
`

type User struct {
	ID    	int    `json:"id"`
	Hobbies []string `json:"hobbies"`
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
	hobbies := []string{"hiking", "biking"}
	err = db.QueryRowContext(ctx, insert, pq.Array(&hobbies)).Scan(&id)
	if err != nil {
		panic(err)
	}


	// update data (works)
	newHobbies := []string{"swimming", "running"}
	u, err := UpdateUser(ctx, db, id, newHobbies) 
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Updated: %v\n", u)
}

func UpdateUser(ctx context.Context, db *sql.DB, id int64, newHobbies []string) (User, error) {
	row := db.QueryRowContext(ctx, update, id, pq.Array(&newHobbies))
	var u User
	err := row.Scan(&u.ID, pq.Array(&u.Hobbies))
	return u, err
}