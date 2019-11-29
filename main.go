package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/gql", handler)
	db.openJSON()
	log.Fatal(http.ListenAndServe(":7777", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HI welcome to the servder")
}

// create the structs to hold the loan
type Loan struct {
	Amount    int       `json:"amount"`
	StartDate time.Time `json:"startDate"`
	Rate      float32   `json:"rate"`
}

type Loans struct {
	Loans []Loan `json:"loans"`
}
