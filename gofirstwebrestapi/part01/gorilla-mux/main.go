package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// declare object router
	router := mux.NewRouter()

	// handler http
	router.HandleFunc("/category/{id:[0-9]+}", categoryHandler) //id itu REGEX

	http.Handle("/", router)
	http.ListenAndServe(":8888", nil)

}

func categoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryId := vars["id"]

	fileName := categoryId + ".html"

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		log.Printf("No such category id")
		fileName = "invalid.html"
	}

	http.ServeFile(w, r, fileName)
}
