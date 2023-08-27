package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

func IsReachable(ctx context.Context, urls []string) error {
	var errs error
	ticker := time.NewTicker(1 * time.Minute)
	for {
		select {
		case <- ticker.C:
			fmt.Println("ticked >>>")
			for _, url := range urls {
				resp, err := http.Head(url)
				if err != nil {
					err = fmt.Errorf("%s : %s", url, err.Error())
					errs = errors.Join(errs, err)
					continue
				}
				defer resp.Body.Close()
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					err = fmt.Errorf("%s : %s", url, err.Error())
					errs = errors.Join(errs, err)
					continue
				}
				fmt.Printf(string(body))  
			}
		case <- ctx.Done():
					
		}

	}
}