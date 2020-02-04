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
		// get the loan struct
		loan := database.Loan{}

		// decode the body into the loan type
		err := json.NewDecoder(r.Body).Decode(&loan)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		// add to the database
		js, err := json.Marshal(database.Create(loan))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// return the loan that was created
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	case "DELETE":
		// get the loan struct
		loan := database.Loan{}

		// decode the body into the loan type
		err := json.NewDecoder(r.Body).Decode(&loan)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		// add to the database
		l, err := database.Delete(loan.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		js, err := json.Marshal(l)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// return the loan that was created
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
