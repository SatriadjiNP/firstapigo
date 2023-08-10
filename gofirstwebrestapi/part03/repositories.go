package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Find all category and return as JSON
func findAllCategories() []byte { //byte untuk memudahkan dan ringan ketika dikirimkan ke web
	var categories Categories
	var category Category

	CategoryResults, err := database.Query("select category_id, category_name, description from categories")
	if err != nil {
		log.Fatal(err)
	}
	defer CategoryResults.Close()

	for CategoryResults.Next() {
		CategoryResults.Scan(&category.CategoryId, &category.CategoryName, &category.Description)
		categories = append(categories, category)
	}

	jsonCategories, err := json.Marshal(categories)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	return jsonCategories
}

func findCategoryById(id int) []byte {
	var category Category

	//fetch one record
	err := database.QueryRow(`
	select category_id, category_name, description 
	from categories where category_id =$1`, id).Scan(&category.CategoryId, &category.CategoryName, &category.Description)
	if err != nil {
		log.Fatal(err)
	}

	jsonCategory, err := json.Marshal(category)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	return jsonCategory
}

func AddCategory(category Category) []byte {

	var addResult ResponseMsg

	// Create prepared statement
	stmt, err := database.Prepare(`
	INSERT INTO categories(category_id,category_name,description) VALUES($1,$2,$3)`)
	if err != nil {
		log.Fatal(err)
	}

	// Execute the prepared statement and retrieve the results
	res, err := stmt.Exec(category.CategoryId, category.CategoryName, category.Description)
	if err != nil {
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	// Populate DBUpdate struct with last Id and num rows affected
	addResult.Affected = rowCnt

	// Convert to JSON and return
	newCategory, err := json.Marshal(addResult)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	return newCategory
}

func deleteCategory(id int64) []byte {
	var deleteResult ResponseMsg

	// Create prepared statement
	stmt, err := database.Prepare(`DELETE FROM categories WHERE category_id=$1`)
	if err != nil {
		log.Fatal(err)
	}

	// Execute the prepared statement and retrieve the results
	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	// Populate DBUpdate struct with last Id and num rows affected
	deleteResult.Id = id
	deleteResult.Affected = rowCnt

	// Convert to JSON and return
	deletedCategory, err := json.Marshal(deleteResult)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	return deletedCategory
}

func updateCategory(category Category) []byte {
	var updateResult ResponseMsg

	// Create prepared statement
	stmt, err := database.Prepare(`
	UPDATE categories set category_name=$1, description=$2
	WHERE category_id=$3
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Execute the prepared statement and retrieve the results
	res, err := stmt.Exec(category.CategoryName, category.Description, category.CategoryId)
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	// Populate DBUpdate struct with last Id and num rows affected
	updateResult.Id = int64(category.CategoryId)
	updateResult.Affected = rowCnt

	// Convert to JSON and return
	updateCategory, err := json.Marshal(updateResult)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	return updateCategory
}
