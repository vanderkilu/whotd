package main

import (
	"log"

	"github.com/vanderkilu/whotd/whotd"
)

func main() {
	event := whotd.NewEvent()
	eventResponse, err := event.CrawlDatePage()
	if err != nil {
		log.Fatal(err)
	}
	event.FormatResponse(eventResponse)
}
