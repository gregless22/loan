package models

import (
	"fmt"
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
func (l Loan) init() {
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

// Create saves the loan to the database
func (l Loan) Create() {
	// change into a valid sqlstatement
	db := database.DB()

	var id int64

	sqlStatement := `INSERT INTO loans (LOAN_FROM, LOAN_TO, START_DATE, END_DATE, RATE, AMOUNT) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := db.QueryRow(sqlStatement, l.From, l.To, l.StartDate, l.EndDate, l.Rate, l.Amount).Scan(&id)

	fmt.Println(id)

	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("saving to the database")
}
