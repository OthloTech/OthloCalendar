package main

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"time"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	connpassEndpoint = "https://connpass.com/api/v1/event/"
	atndEndpoint     = "http://api.atnd.org/events/"
)

func main() {
	now := time.Now()
	//ym := fmt.Sprintf("%v%v", now.Year(), int(now.Month()))

	query := Query{Start: 1, Order: CREATE}
	query.KeywordOr = strings.Split("名古屋,python", ",")
	query.Time = []Time{Time{Year: now.Year(), Month: int(now.Month())}}
	query.Count = 5

	var events []Events

	res, err := query.Connpass()
	if err != nil {
		fmt.Errorf("Failed to execute search: %v.", err)
		return
	}
	events = res.Events

	for _, e := range events {
		spew.Dump(e)
	}
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
// func connpass(keyword, ym string) ([]Events, error) {
// 	url := fmt.Sprintf("%s?keyword=%s&ym=%s&count=10", connpassEndpoint, keyword, ym)
// 	res, err := http.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	var eventData EventData

// 	err = json.Unmarshal(body, &eventData)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return eventData.Events, nil
// }

// func atnd(keyword, ym string) ([]Events, error) {
// 	url := fmt.Sprintf("%s?keyword=%s&ym=%s&count=10", atndEndpoint, keyword, ym)
// 	res, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	var eventData EventData

// 	//escapedBody := []byte(strings.Replace(body, "'", "\\'", -1))

// 	err = json.Unmarshal(body, &eventData)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}
// 	fmt.Println(eventData.ResultsReturned)
// 	return eventData.Events, nil
// }

type EventData struct {
	ResultsReturned  int      `json:"results_returned"`
	ResultsStart     int      `json:"results_start"`
	ResultsAvailable int      `json:"results_available"`
	Events           []Events `json:"events"`
}

type Events struct {
	EventId     int    `json:"event_id"`
	EventURL    string `json:"event_url"`
	EventType   string `json:"event_type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Catch       string `json:"catch"`
	HashTag     string `json:"hash_tag"`
	Accepted    int    `json:"accepted"`
	Waiting     int    `json:"waiting"`
	Limit       int    `json:"limit"`
	URL         string `json:"url"`

	OwnerID          int    `json:"owner_id"`
	OwnerNickname    string `json:"owner_nickname"`
	OwnerTwitterID   string `json:"owner_twitter_id"`
	OwnerDisplayName string `json:"owner_display_name"`

	UpdatedAt time.Time `json:"updated_at"`
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at"`

	Address string `json:"address"`
	Place   string `json:"place"`
	Lat     string `json:"lat"`
	Lon     string `json:"lon"`

	Series Series `json:"series"`
}

type Series struct {
	URL   string `json:"url"`
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Order int
const (
	UPDATE Order = 1 + iota // 1: descending in updated time
	START                   // 2: descending in event start time
	CREATE                  // 3: descending in created time
)

type Time struct {
	Year  int
	Month int
	Date  int
}

type Format string
const (
	JSON Format = "json"
)

type Query struct {
	EventId     []int
	KeywordAnd  []string
	KeywordOr   []string
	Time        []Time
	Participant []string
	Owner       []string
	SeriesId    []int
	Start       int
	Order       Order
	Count       int
	Format
}

func (q Query) buildURL(baseUrl string) string {
	u, err := url.Parse(baseUrl)
	if err != nil {
		log.Fatal(err)
	}
	v := url.Values{}
	setInts(v, "event_id", q.EventId)
	setStrings(v, "keyword", q.KeywordAnd)
	setStrings(v, "keyword_or", q.KeywordOr)
	if q.Time != nil && len(q.Time) > 0 {
		ymd, ym := printTimeArray(q.Time)
		if len(ymd) > 0 {
			v.Set("ymd", ymd)
		}
		if len(ym) > 0 {
			v.Set("ym", ym)
		}
	}
	setStrings(v, "nickname", q.Participant)
	setStrings(v, "owner_nickname", q.Owner)
	if q.Start > 0 {
		v.Set("start", fmt.Sprint(q.Start))
	}
	if q.Order > 0 {
		v.Set("order", fmt.Sprint(q.Order))
	}
	if q.Count > 0 {
		v.Set("count", fmt.Sprint(q.Count))
	}
	u.RawQuery = v.Encode()
	return u.String()
}

func setInts(p url.Values, k string, v []int) {
	if v != nil {
		for _, n := range v {
			p.Add(k, strconv.Itoa(n))
		}
	}
}

func setStrings(p url.Values, k string, v []string) {
	if v != nil {
		for _, e := range v {
			p.Add(k, e)
		}
	}
}

func printTimeArray(arr []Time) (string, string) {
	ymd := make([]string, 0)
	ym := make([]string, 0)
	for _, v := range arr {
		if v.Year > 0 && v.Month > 0 {
			if v.Date > 0 {
				ymd = append(ymd, fmt.Sprintf("%04d%02d%02d", v.Year, v.Month, v.Date))
			} else {
				ym = append(ym, fmt.Sprintf("%04d%02d", v.Year, v.Month))
			}
		}
	}
	return strings.Join(ymd, ","), strings.Join(ym, ",")
}

func parse(jsonBlob []byte) (*EventData, error) {
	res := new(EventData)
	err := json.Unmarshal(jsonBlob, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (q Query) Connpass() (*EventData, error) {
	url := q.buildURL(connpassEndpoint)
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	return parse(body)
}
