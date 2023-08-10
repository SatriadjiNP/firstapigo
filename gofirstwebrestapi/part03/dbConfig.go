package main

import (
	"database/sql"
	"log"
)

var database *sql.DB

func dbConnect() {
	// declare connection string
	connStr := "user=postgres password=admin dbname=NorthwindDB sslmode=disable port=5432"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	database = db
	log.Println("DB Connected")
}
