package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	dbConnect()
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8888", router))
}
