package main

type Category struct {
	CategoryId   int    `json:"id"`
	CategoryName string `json:"category"`
	Description  string `json:"description"`
}

type Categories []Category
