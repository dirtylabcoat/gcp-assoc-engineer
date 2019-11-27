package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dateTime := dateTimeAsString()
	out := fmt.Sprintf("{\"dateTime\":\"%s\"}", dateTime)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(out))
}

func dateTimeAsString() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02 15:04:05 Monday")
}

func main() {
	s := &server{}
	http.Handle("/datetime", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

