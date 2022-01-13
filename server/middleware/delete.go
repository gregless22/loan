package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gregless22/loan/database"
)

// DeleteLoan will delete the loan from the database
func DeleteLoan(w http.ResponseWriter, r *http.Request) {
	// get the ID from the request URL
	split := strings.Split(r.URL.Path, "/")
	url := split[len(split)-1]

	//check is an int
	id, err := strconv.ParseInt(url, 10, 64)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintf(w, "Not a number in URL")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deleteLoan(id))
}

func deleteLoan(id int64) int64 {

	// create the postgres db connection
	db := database.Connect()

	// close the db connection
	defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM loans WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected
}
