package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)



type Indexer struct {
	duration  time.Duration
	urls      []string
}


func New(d time.Duration, urls []string) Indexer {
	return Indexer{	
		duration: d,
		urls: urls,
	}
}

func (i *Indexer) AppendNewUrls(urls []string) {
	i.urls = append(i.urls, urls...)
}

func (i *Indexer) Start(stop <-chan os.Signal, done chan int) {
	ticker := time.NewTicker(i.duration)

	// Run the cron job loop
	for {
		select {
		case <-ticker.C:
			fmt.Println("ticked")
			n := len(i.urls)
			for k, url := range i.urls {
				fmt.Printf("url %d processed: %s\n", k, url)
				time.Sleep(100 * time.Millisecond)
			}
			if len(i.urls) > n {
				for k := n; n < len(i.urls); k++ {
					fmt.Printf("url %d processed: %s\n", k, i.urls[k])
				}
			}
		case <-stop:
			fmt.Println("stop signal received")
			ticker.Stop()
			done <- 1
		}
	}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	

	urls := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	i := New(5 * time.Second, urls)
	go func(){
		if err := i.StartWithContext(ctx); err != nil {
			fmt.Println(err)
		}
	}()
}

func (i *Indexer) StartWithContext(ctx context.Context) error {
	ticker := time.NewTicker(i.duration)

	// Run the cron job loop
	for {
		select {
		case <-ticker.C:
			fmt.Println("ticked")
			n := len(i.urls)
			for k, url := range i.urls {
				fmt.Printf("url %d processed: %s\n", k, url)
				time.Sleep(100 * time.Millisecond)
			}
			if len(i.urls) > n {
				for k := n; n < len(i.urls); k++ {
					fmt.Printf("url %d processed: %s\n", k, i.urls[k])
				}
			}
			return nil
		case <- ctx.Done():
			fmt.Println("stop signal received")
			ticker.Stop()
			return ctx.Err()
		}
	}
}
