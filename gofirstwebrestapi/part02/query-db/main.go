package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Category struct {
	CategoryId   int
	CategoryName string
	Description  string
}

var database *sql.DB

func main() {
	// declare connection string
	connStr := "user=postgres password=admin dbname=NorthwindDB sslmode=disable port=5432"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	database = db
	log.Println("DB Connected")

	http.HandleFunc("/", showCategory)
	http.HandleFunc("/category", displayCategory)

	log.Println("Starting server on 8888")

	errHttp := http.ListenAndServe(":8888", nil)
	if errHttp != nil {
		log.Println(errHttp)
	}
}

func showCategory(w http.ResponseWriter, r *http.Request) {
	category := Category{}

	rows, err := database.Query(`
		select category_id, category_name,
		description from categories
	`)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// tampilkan output row category
	for rows.Next() {
		err := rows.Scan(&category.CategoryId, &category.CategoryName, &category.Description)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "%d : %s : %s \n", category.CategoryId, category.CategoryName, category.Description)
	}
}

func displayCategory(w http.ResponseWriter, r *http.Request) {
	categories := []Category{}

	rows, err := database.Query(`
		select category_id, category_name,
		description from categories
	`)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// tampilkan output row category
	for rows.Next() {
		category := Category{}
		err := rows.Scan(&category.CategoryId, &category.CategoryName, &category.Description)
		if err != nil {
			log.Fatal(err)
		}
		categories = append(categories, category)
	}
	// load category.html for manipulation with data from slice
	t, _ := template.ParseFiles("category.html")

	t.Execute(w, categories)
}
