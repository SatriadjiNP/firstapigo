package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func findAllCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	// Query the database
	jsonCategories := findAllCategories()

	// Format the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonCategories)
}

func findCategoryHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	categoryId, _ := strconv.Atoi(vars["id"])

	// Query the database
	jsonCategory := findCategoryById(categoryId)

	// Format the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonCategory)
}

func AddCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var category Category

	// Read the body of the request
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// Convert the JSON in the request to a Go type
	if err := json.Unmarshal(body, &category); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	// Write to the database
	addResult := AddCategory(category)

	// Format the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(addResult)
}

func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {

	// Get URL parameter with the category ID to delete
	vars := mux.Vars(r)
	categoryId, _ := strconv.ParseInt(vars["id"], 10, 64)

	// Query the database
	deleteResult := deleteCategory(categoryId)

	// Format the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(deleteResult)
}

func UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var category Category

	// Read the body of the request
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// Convert the JSON in the request to a Go type
	if err := json.Unmarshal(body, &category); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	// Write to the database
	updateResult := updateCategory(category)

	// Format the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(updateResult)
}
