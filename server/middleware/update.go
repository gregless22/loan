package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gregless22/loan/database"
	"github.com/gregless22/loan/models"
)

// UpdateLoan will take a loan type and update it in the database
func UpdateLoan(w http.ResponseWriter, r *http.Request) {
	var loan models.Loan

	err := json.NewDecoder(r.Body).Decode(&loan)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding data in PUT: %s", err)
		return
	}

	updateLoan(loan)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Loan Sucessfully updated")
}

func updateLoan(l models.Loan) {
	// connect to database
	db := database.Connect()
	defer db.Close()

	// creat the sql statement
	sqlStatement := `UPDATE loans SET LOAN_FROM=$2, LOAN_TO=$3, START_DATE=$4, END_DATE=$5, RATE=$6, AMOUNT=$7 WHERE ID=$1`

	_, err := db.Exec(sqlStatement, l.ID, l.From, l.To, l.StartDate, l.EndDate, l.Rate, l.Amount)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

}
