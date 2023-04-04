package main


import (
	"github.com/mmcdole/gofeed"
)

func main(){
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://feeds.transistor.fm/cup-o-go")
	for _, item := range feed.Items {
		println(item.Title)
	}
}