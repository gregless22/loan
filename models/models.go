package models

import "time"

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

// InitLoan will return the psql command to insert in to the database
func InitLoan() string {
	return `CREATE TABLE if NOT EXISTS loan (
		id          SMALLSERIAL   PRIMARY KEY,
		from        CHAR(30)      NOT NULL,
		to          CHAR(30)      NOT NULL,
    start_date  DATE          NOT NULL,
    end_date    DATE,
    rate        REAL          NOT NULL,
    amount      MONEY         NOT NULL )
		`
}
