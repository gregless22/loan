package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// Loan holds the loan struct
type Loan struct {
	Amount    int       `json:"amount"`
	StartDate time.Time `json:"startDate"`
	Rate      float32   `json:"rate"`
}

// Loans holds a slice of Loans
type Loans struct {
	Loans []Loan `json:"loans"`
}

// OpenJSON opens the file read to write
func openJSON() (*os.File, error) {
	jsonFile, err := os.Open("loan.json")

	// handle the err
	if err != nil {
		fmt.Println(err)
	}

	return jsonFile, err

}

// Handler is just a test
func Handler() {
	openJSON()
}

// Read will return all of the loans
func Read() Loans {
	jsonFile, err := openJSON()
	defer jsonFile.Close()

	// Read the JSON
	byte, _ := ioutil.ReadAll(jsonFile)

	//initialise the loans array
	var loans Loans

	// unmarshal the bytes into the struct Loan
	err = json.Unmarshal(byte, &loans)
	if err != nil {
		fmt.Println(err)
	}

	// return the loans
	return loans
}
