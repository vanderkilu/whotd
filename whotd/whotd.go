package whotd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
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
		//skip first content as it is not important
		//and also target the next 3 content only
		if index <= 3 && index != 0 {
			selection.Find("li").Each(func(i int, s *goquery.Selection) {
				if index == 1 {
					events = append(events, s.Text())
				} else if index == 2 {
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
	color.Blue("PROMINENTS EVENTS THAT OCCURED ON THIS DAY")
	fmt.Println()

	for _, event := range response.events {
		fmt.Printf("%s\n\n", event)
	}

	color.Green("BIRTHS OF PROMINENT PEOPLE ON THIS DAY")
	fmt.Println()

	for _, event := range response.births {
		fmt.Printf("%s\n\n", event)
	}

	color.Red("DEATHS OF PROMINENT PEOPLE ON THIS DAY\n")
	fmt.Println()

	for _, event := range response.deaths {
		fmt.Printf("%s\n\n", event)
	}

}
