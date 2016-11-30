package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/davecgh/go-spew/spew"
	"github.com/OthloTech/OthloCalendar/server/models"
)

const (
	connpassEndpoint = "https://connpass.com/api/v1/event/"
	atndEndpoint     = "http://api.atnd.org/events/"
	zusaarEndpoint   = "http://www.zusaar.com/api/event/"
)

type EventController struct{}

var eventModel = new(models.EventModel)

func (ctrl EventController) Search(c *gin.Context) {
	query := Query{Start: 1, Order: CREATE}
	if k := c.Query("keyword"); len(k) != 0 {
		query.KeywordAnd = strings.Split(k, ",")
	}
	if k := c.Query("keyword_or"); len(k) != 0 {
		query.KeywordOr = strings.Split(k, ",")
	}
	if o := c.Query("owner"); len(o) != 0 {
		query.Owner = strings.Split(o, ",")
	}
	if p := c.Query("participant"); len(p) != 0 {
		query.Participant = strings.Split(p, ",")
	}
	//if i := c.Query("event_id"); len(i) != 0 {
	//	query.EventId = strings.Split(i, ",")
	//}

	now := time.Now()
	query.Time = []Time{Time{Year: now.Year(), Month: int(now.Month())}}
	query.Count = 100

	var events []Events

	res, err := query.Connpass()
	if err != nil {
		fmt.Errorf("Failed to execute search: %v.", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		return
	}
	events = res.Events

	for _, e := range events {
		spew.Dump(e)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event found", "events": events})


	// for {
	// 	res, err := query.Search()
	// 	if err != nil {
	// 		fmt.Errorf("Failed to execute search: %v.", err)
	// 		return
	// 	}
	// 	allEvents = append(allEvents, res.Events...)
	// 	offset := res.Start + res.Returned
	// 	if offset > res.Available {
	// 		break
	// 	}
	// 	query.Start = offset
	// }
}

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
