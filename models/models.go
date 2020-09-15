package models

import (
	"time"

	"github.com/gregless22/loan/database"
)

// Loan schema for storing in the database
type Loan struct {
	ID        int64     `json:"id"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Rate      float32   `json:"rate"`
	Amount    int64     `json:"amount"`
}

// Init will initiallise the database table for the loan
func (l Loan) Init() {
	sqlStatement := `CREATE TABLE if NOT EXISTS LOANS (
		LOAN_FROM        CHAR(30)      NOT NULL,
		LOAN_TO          CHAR(30)      NOT NULL,
    START_DATE  DATE          NOT NULL,
    END_DATE    DATE,
    RATE        REAL          NOT NULL,
    AMOUNT      MONEY         NOT NULL, 
    ID          SMALLSERIAL        PRIMARY KEY
  )`

	database.Init(sqlStatement)
}

// GetAll returns all of the loans in the database
// func (l Loan) GetAll() ([]Loan, error) {
// 	//create empty slice of loans
// 	var loans []Loan

// 	// connect to the database
// 	db := database.DB()
// 	defer db.Close()

// 	// create the select sql query
// 	sqlStatement := `SELECT * FROM loans`

// 	// execute the sql statement
// 	rows, err := db.Query(sqlStatement)
// 	defer rows.Close()

// 	if err != nil {
// 		log.Fatalf("Unable to execute the query. %v", err)
// 	}

// 	// iterate over the rows
// 	for rows.Next() {
// 		var loan Loan

// 		// unmarshal the row object to user
// 		err = rows.Scan(&loan.From, &loan.To, &loan.StartDate, &loan.EndDate, &loan.Rate, &loan.Amount, &loan.ID)

// 		if err != nil {
// 			log.Fatalf("Unable to scan the row. %v", err)
// 		}

// 		// append the user in the users slice
// 		loans = append(loans, loan)

// 	}

// 	// return empty user on error
// 	return loans, err

// }
