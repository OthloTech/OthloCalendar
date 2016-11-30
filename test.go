package main

import (
	//"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"fmt"
)

var apis map[string]string = map[string]string {
	"connpass": "https://connpass.com/api/v1/event/",
}

func main() {
	now := time.Now()
	ym := fmt.Sprintf("%v%v", now.Year(), int(now.Month()))
	fmt.Println(ym)

	fmt.Println(apis["connpass"])

	resp, err := http.Get("https://connpass.com/api/v1/event/?keyword=python&count=1")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(string(body))
}

func (keyword string) Connpass() {
	
}