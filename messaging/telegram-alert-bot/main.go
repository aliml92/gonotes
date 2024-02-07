package main

import (
	"log"

	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	pref := tele.Settings{
		Token:  "<REDACTED>",
		Poller: &tele.LongPoller{Timeout: 5 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}


	group := tele.ChatID(-1000000000000)

	b.Send(group, "test", &tele.SendOptions{})

}

