package main

import (
	"log"
	"net/http"
	"time"
)
type Feed struct {
	Url string
	LastModified time.Time
	ETag string
	Age time.Duration
}


func main(){
	var feeds []*Feed
	feeds = append(feeds, &Feed{
		Url: "https://feeds.transistor.fm/cup-o-go",
	})
	feeds = append(feeds, &Feed{
		Url: "https://thedotnetcorepodcast.libsyn.com/rss",
	})
	// run first time
	Run(feeds)
	// run second time
	Run(feeds)
}

func Run(feeds []*Feed){
	for _, feed := range feeds {
		log.Printf("Checking feed %v\n", feed)
		req, err := http.NewRequest("GET", feed.Url, nil)
		if err != nil {
			log.Fatalf("Error creating request for %s: %v", feed.Url, err)
		}
		if !feed.LastModified.IsZero() {
			req.Header.Set("If-Modified-Since", feed.LastModified.Format(http.TimeFormat))
		}
		if feed.ETag != "" {
			req.Header.Set("If-None-Match", feed.ETag)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatalf("Error fetching %s: %v", feed.Url, err)
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusNotModified {
			log.Printf("Feed %s not modified", feed.Url)
			continue
		}
		if resp.StatusCode != http.StatusOK {
			log.Printf("Feed %s returned status %d", feed.Url, resp.StatusCode)
			continue
		}
		feed.LastModified, err = time.Parse(http.TimeFormat, resp.Header.Get("Last-Modified"))
		if err != nil {
			log.Fatalf("Error parsing last modified time for %s: %v", feed.Url, err)
		}
		feed.ETag = resp.Header.Get("ETag")
		log.Printf("Feed %s modified", feed.Url)
		log.Printf("Last modified: %s", feed.LastModified)
		log.Printf("ETag: %s", feed.ETag)
		resp.Header.Write(log.Writer())
	}
}