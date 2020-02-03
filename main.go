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
	http.HandleFunc("/loan", loan)

	log.Fatal(http.ListenAndServe(":7777", nil))
}

// TODO Update

// TODO Delete

func loan(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/loan" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		js, err := json.Marshal(database.Read())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	case "POST":
		fmt.Fprint(w, "This works")

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
