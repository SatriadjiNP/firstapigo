package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var database *sql.DB

func main() {
	// declare connection string
	connStr := "user=postgres dbname=NorthwindDB sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	database = db
	log.Println("DB Connected")
}
