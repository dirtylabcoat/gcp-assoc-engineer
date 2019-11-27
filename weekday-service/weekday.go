package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{\"weekday\":\"%s\"}", time.Now().Format("Monday"))))
}

func main() {
	s := &server{}
	http.Handle("/weekday", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
