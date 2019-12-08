package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type server struct{}

type Quote struct {
	Id     int
	Author string
	Quote  string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbHost := "10.57.32.3"
	dbPort := "3306"
	dbUser := "root"
	dbPass := "testpass1"
	dbName := "ace_exam_book"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	quote := getQuote()
	out := fmt.Sprintf("{\"quote\":\"%s\"}", quote)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(out))
}

func getQuote() string {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM quote")
	if err != nil {
		panic(err.Error())
	}
	q := Quote{}
	res := []Quote{}
	for selDB.Next() {
		var id int
		var author, quote string
		err = selDB.Scan(&id, &author, &quote)
		if err != nil {
			panic(err.Error())
		}
		q.Id = id
		q.Author = author
		q.Quote = quote
		res = append(res, q)
	}
	defer db.Close()
	return res[0].Quote
}

func main() {
	s := &server{}
	http.Handle("/quote", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
