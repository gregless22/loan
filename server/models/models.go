package models

import (
	"time"

	"github.com/gregless22/loan/server/database"
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
