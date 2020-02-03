package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gregless22/loan/database"
)

func main() {
	// TODO import the handlers
	http.HandleFunc("/test", handler)
	http.HandleFunc("/create", create)
	http.HandleFunc("/read", read)

	log.Fatal(http.ListenAndServe(":7777", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// test the database
	database.Handler()
}

func create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this will eventually write")
}

// TODO Read
func read(w http.ResponseWriter, r *http.Request) {
	// get the data
	js, err := json.Marshal(database.Read())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// TODO Update

// TODO Delete
