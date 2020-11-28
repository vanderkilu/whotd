package whotd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var baseURL = "https://en.wikipedia.org/wiki/"

type Event struct {
	date      time.Time
	eventType []string
}

type EventResponse struct {
	events []string
	births []string
	deaths []string
}

func NewEvent() *Event {
	events := []string{"events", "births", "deaths"}
	return &Event{date: time.Now(), eventType: events}
}

func (e *Event) getMonthStr() string {
	_, month, day := e.date.Date()
	return fmt.Sprintf("%v_%d", month, day)

}

func (e *Event) CrawlDatePage() (EventResponse, error) {
	var eventResponse EventResponse
	url := fmt.Sprintf("%s%s", baseURL, e.getMonthStr())
	resp, err := http.Get(url)
	if err != nil {
		return eventResponse, err
	}
	defer resp.Body.Close()
	document, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return eventResponse, err
	}
	selector := ".mw-parser-output ul"
	var events []string
	var births []string
	var deaths []string
	document.Find(selector).Each(func(index int, selection *goquery.Selection) {
		if index <= 2 {
			selection.Find("li").Each(func(i int, s *goquery.Selection) {
				if index == 0 {
					events = append(events, s.Text())
				} else if index == 1 {
					births = append(births, s.Text())
				} else {
					deaths = append(deaths, s.Text())
				}
			})
		}
		eventResponse = EventResponse{events, births, deaths}
	})
	return eventResponse, nil
}

func (e *Event) FormatResponse(response EventResponse) {
	for _, event := range response.events {
		fmt.Println(event)
	}

	fmt.Println()
	fmt.Println(".....................................")
	fmt.Println()

	for _, event := range response.births {
		fmt.Println(event)
	}

	fmt.Println()
	fmt.Println(".....................................")
	fmt.Println()

	for _, event := range response.deaths {
		fmt.Println(event)
	}

}
