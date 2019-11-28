package main

import (
    "encoding/json"
	"fmt"
    "io/ioutil"
	"log"
	"net/http"
	"time"
)

type server struct{}

type Response struct {
    Weekday string `json:"weekday"`
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dateTime := dateTimeAsString()
    weekday := getWeekday() 
	out := fmt.Sprintf("{\"dateTime\":\"%s %s\"}", dateTime, weekday)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(out))
}

func dateTimeAsString() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02 15:04:05")
}

func getWeekday() string {
    response, err := http.Get("http://weekday.default:8080/weekday")
    if err != nil {
        return "FAIL"
    } else {
        var res Response
        data, _ := ioutil.ReadAll(response.Body)
        json.Unmarshal(data, &res)
        return res.Weekday
    }
}

func main() {
	s := &server{}
	http.Handle("/datetime", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

