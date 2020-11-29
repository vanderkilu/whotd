package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/vanderkilu/whotd/whotd"
)

//check if date string provided is valid
//if valid convert to query format(eg November_11) and return it else throw an error.
func formateDate(dateStr string) (string, error) {
	dateParts := strings.Split(dateStr, "/")
	if len(dateParts) != 2 {
		return "", errors.New("Invalid date provided")
	}
	fullDateStr := fmt.Sprintf("2006-%s-%s", dateParts[1], dateParts[0])
	t, err := time.Parse("2006-01-02", fullDateStr)
	if err != nil {
		return "", errors.New("Invalid date provided. Must be in format dd/mm Example 08/11")
	}
	_, month, day := t.Date()
	query := fmt.Sprintf("%v_%d", month, day)
	return query, nil

}

func main() {
	var event *whotd.Event
	date := flag.String("d", "", "date to search events for. Must be in format dd/mm Example 08/11, 05/05 etc ")
	flag.Parse()

	if *date != "" {
		query, err := formateDate(*date)
		if err != nil {
			log.Fatal(err)
		}
		event = whotd.NewEvent(query)
	} else {
		_, month, day := time.Now().Date()
		query := fmt.Sprintf("%v_%d", month, day)
		event = whotd.NewEvent(query)
	}

	fmt.Println("Fetching Events......")

	eventResponse, err := event.CrawlDatePage()
	if err != nil {
		log.Fatal(err)
	}
	event.FormatResponse(eventResponse)
}
