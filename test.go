package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	connpassEndpoint = "https://connpass.com/api/v1/event/"
)

func main() {
	now := time.Now()
	ym := fmt.Sprintf("%v%v", now.Year(), int(now.Month()))

	events := connpass("python", ym)
	for i, event := range events {
  		fmt.Print(fmt.Sprintf("index:%d,value:%s\n", i, event.Title))
	}
}

func connpass(keyword, ym string) []Events {
	url := fmt.Sprintf("%s?keyword=%s&ym=%s&count=10", connpassEndpoint, keyword, ym)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	var eventData EventData

	err = json.Unmarshal(body, &eventData)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(eventData.ResultsReturned)
	return eventData.Events
}

type EventData struct {
	ResultsReturned int
	Events []Events
	ResultsStart int
	ResultsAvailable int
}

type Events struct {
	EventURL string
	EventType string
	OwnerNickname string
	Series Series
	UpdatedAt time.Time
	Lat string
	StartedAt time.Time
	HashTag string
	Title string
	EventId int
	Lon string
	Waiting int
	Limit int
	OwnerID int
	OwnerDisplayName string
	Description string
	Address string
	EndedAt time.Time
	Place string
}

type Series struct {
	URL string
	ID int
	Title string
}
