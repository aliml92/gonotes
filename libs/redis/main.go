package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)


func Connect() (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		MinIdleConns: 200,
		PoolSize:     1200,
		PoolTimeout:  time.Duration(240) * time.Second,
		Username:     "test_user",
		Password:     "sample_pass", 
		DB:           0,       
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func main() {
	c, err := Connect()
	if err != nil {
		fmt.Printf("error occured: %s\n", err.Error())
	}
	fmt.Printf("client: %v\n", c)
}