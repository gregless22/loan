package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// Loan holds the loan struct
type Loan struct {
	ID        int       `json:"id"`
	Amount    int       `json:"amount"`
	StartDate time.Time `json:"startDate"`
	Rate      float32   `json:"rate"`
}

// Loans holds a slice of Loans
type Loans struct {
	Loans []Loan `json:"loans"`
}

// OpenJSON opens the file read to write
func openJSON() *os.File {
	jsonFile, err := os.Open("loan.json")

	// handle the err
	if err != nil {
		fmt.Println(err)
	}

	return jsonFile

}

// Read will return all of the loans
func Read() Loans {
	jsonFile := openJSON()
	defer jsonFile.Close()

	// Read the JSON
	byte, _ := ioutil.ReadAll(jsonFile)

	//initialise the loans array
	var loans Loans

	// unmarshal the bytes into the struct Loan
	err := json.Unmarshal(byte, &loans)
	if err != nil {
		fmt.Println(err)
	}

	// return the loans
	return loans
}

// Create will create a new instance of loan.  It recieves a Loan
func Create(l Loan) Loan {
	loans := Read()

	index := 1
	// TODO add the loan ID
	for _, loan := range loans.Loans {
		if loan.ID >= index {
			index = loan.ID + 1
		}
	}
	l.ID = index

	// add to the array
	loans.Loans = append(loans.Loans, l)

	// rewrite to the json file
	byte, err := json.Marshal(loans)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("loan.json", byte, 0644)
	if err != nil {
		fmt.Println(err)
	}
	return l

}

// Delete will remove the loan
func Delete(id int) (Loan, error) {
	loans := Read()

	l := Loan{}
	// find the Loan
	for i, loan := range loans.Loans {
		fmt.Println(id)
		fmt.Println(loan.ID)
		if id == loan.ID {
			l = loan
			loans.Loans = append(loans.Loans[:i], loans.Loans[i+1:]...)
			break
		}
		if i == len(loans.Loans) {
			//return not found
			err := errors.New("Unable to find current loan in database")
			return l, err
		}

	}

	// rewrite to the json file
	byte, err := json.Marshal(loans)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("loan.json", byte, 0644)
	if err != nil {
		fmt.Println(err)
	}
	return l, err
}
