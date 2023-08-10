package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// declare http handler
	// default servermux
	http.HandleFunc("/", welcome)

	log.Println("Server port listening on port : 8888")

	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		log.Fatal("Listener Error : ", err)
	}
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Current time : ", time.Now())
}
