package models

// Loan is the struct for one loan object
type Loan struct {
	Amount    int       `json:"amount"`
	StartDate time.Time `json:"startDate"`
	Rate      float32   `json:"rate"`
}

// Loans holds a slice of Loans
type Loans struct {
	Loans []Loan `json:"loans"`
}


// TODO Create

// TODO Read

// TODO Update

// TODO Delete