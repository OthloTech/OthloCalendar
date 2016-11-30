package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"github.com/davecgh/go-spew/spew"
	//"strings"
)

const (
	connpassEndpoint = "https://connpass.com/api/v1/event/"
	atndEndpoint = "http://api.atnd.org/events/"
)

func main() {
	now := time.Now()
	ym := fmt.Sprintf("%v%v", now.Year(), int(now.Month()))

	events, err := connpass("python", ym)
	if err != nil {
		fmt.Println(err)
	}
	for i, event := range events {
		spew.Dump(event)
		fmt.Println("=======")
		//fmt.Print(fmt.Sprintf("index:%d,value:%s\n", i, event.Title))
	}

	// events, err := atnd("python", ym)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for i, event := range events {
 //  		fmt.Print(fmt.Sprintf("index:%d,value:%s\n", i, event.Title))
	// }	
}

/*
  connpassの使用できるparam See: https://connpass.com/about/api/
  event_id int 
  keyword string
  keyword_or string
  ym string
  nickname string 参加者のニックネーム
  owner_nickname string 管理者のニックネーム
  series_id int
  start int
  order int
  count int
  format string default json
*/
func connpass(keyword, ym string) ([]Events, error) {
	url := fmt.Sprintf("%s?keyword=%s&ym=%s&count=10", connpassEndpoint, keyword, ym)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	var eventData EventData

	err = json.Unmarshal(body, &eventData)
	if err != nil {
		return nil, err
	}
	fmt.Println(eventData.ResultsReturned)
	return eventData.Events, nil
}

func atnd(keyword, ym string) ([]Events, error) {
	url := fmt.Sprintf("%s?keyword=%s&ym=%s&count=10", atndEndpoint, keyword, ym)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	var eventData EventData

	//escapedBody := []byte(strings.Replace(body, "'", "\\'", -1))

	err = json.Unmarshal(body, &eventData)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(eventData.ResultsReturned)
	return eventData.Events, nil
}

type EventData struct {
	ResultsReturned int `json:"results_returned"`
	ResultsStart int `json:"results_start"`
	ResultsAvailable int `json:"results_available"`
	Events []Events `json:"events"`
}

type Events struct {
	EventId int `json:"event_id"`
	EventURL string `json:"event_url"`
	EventType string `json:"event_type"`
	Title string `json:"title"`
	Description string `json:"description"`
	Catch string `json:"catch"`
	HashTag string `json:"hash_tag"`
	Accepted int `json:"accepted"`
	Waiting int `json:"waiting"`
	Limit int `json:"limit"`
	URL string `json:"url"`

	OwnerID int `json:"owner_id"`
	OwnerNickname string `json:"owner_nickname"`
	OwnerTwitterID string `json:"owner_twitter_id"`
	OwnerDisplayName string `json:"owner_display_name"`

	UpdatedAt time.Time `json:"updated_at"`
	StartedAt time.Time `json:"started_at"`
	EndedAt time.Time `json:"ended_at"`

	Address string `json:"address"`
	Place string `json:"place"`
	Lat string `json:"lat"`
	Lon string `json:"lon"`

	Series Series `json:"series"`
}

type Series struct {
	URL string`json:"url"`
	ID int `json:"id"`
	Title string `json:"title"`
}
