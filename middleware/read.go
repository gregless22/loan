package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gregless22/loan/database"
	"github.com/gregless22/loan/models"
)

// ReadLoans will return all of the loans in the database
func ReadLoans(w http.ResponseWriter, r *http.Request) {

	loans, err := readAllLoans()

	if err != nil {
		fmt.Println("Error reading loans from database")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(loans)
}

// get loans from the DB by its userid
func readAllLoans() ([]models.Loan, error) {
	// create the postgres db connection
	db := database.Connect()

	// close the db connection
	defer db.Close()

	var loans []models.Loan

	// create the select sql query
	sqlStatement := `SELECT * FROM loans`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		// create variable for the capture of different types from postgres
		var amount string
		var from string
		var to string

		var l models.Loan

		// unmarshal the row object to user
		err = rows.Scan(&from, &to, &l.StartDate, &l.EndDate, &l.Rate, &amount, &l.ID)

		// convert the results and put into the object
		// remove the "$" sign
		amount = strings.Trim(amount, "$")
		split := strings.Split(amount, ".")
		amount = split[0]
		amount = strings.ReplaceAll(amount, ",", "")
		l.Amount, err = strconv.ParseInt(amount, 10, 64)

		l.From = strings.Trim(from, " ")
		l.To = strings.Trim(to, " ")

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		loans = append(loans, l)

	}

	// return empty user on error
	return loans, err
}
