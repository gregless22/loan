package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gregless22/loan/database"
	"github.com/gregless22/loan/models"
)

// The middleware will complete calls to the back end and encode / decode json

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// CreateLoan will handle a new loan entry
func CreateLoan(w http.ResponseWriter, r *http.Request) {
	// create an emtpy loan
	var loan models.Loan

	// decode json in bytes
	err := json.NewDecoder(r.Body).Decode(&loan)
	if err != nil {
		fmt.Fprintf(w, "Error decoding data: %s", err)
	}

	id := createLoan(loan)

	res := response{
		ID:      id,
		Message: "User created successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(res)
}

func createLoan(l models.Loan) int64 {

	// create the postgres db connection
	db := database.Connect()

	// close the db connection
	defer db.Close()

	// create id that will be filled by the database
	var id int64

	// create the insert sql query
	sqlStatement := `INSERT INTO loans (LOAN_FROM, LOAN_TO, START_DATE, END_DATE, RATE, AMOUNT) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, l.From, l.To, l.StartDate, l.EndDate, l.Rate, l.Amount).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}
