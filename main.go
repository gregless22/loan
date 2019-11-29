package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/gql", handler)
	openJSON()
	log.Fatal(http.ListenAndServe(":7777", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HI welcome to the servder")
}

func openJSON() {
	jsonFile, err := os.Open("loan.json")
	defer jsonFile.Close()
	// handle the err
	if err != nil {
		fmt.Println(err)
	}

	// Read the JSON
	byte, _ := ioutil.ReadAll(jsonFile)

	//initialise the loans array
	var loans Loans

	// unmarshal the bytes into the struct Loan
	err = json.Unmarshal(byte, &loans)
	if err != nil {
		fmt.Println(err)
	}

	//test
	for i := 0; i < len(loans.Loans); i++ {
		fmt.Println(loans.Loans[i])
	}

}

// create the structs to hold the loan
type Loan struct {
	Amount    int       `json:"amount"`
	StartDate time.Time `json:"startDate"`
	Rate      float32   `json:"rate"`
}

type Loans struct {
	Loans []Loan `json:"loans"`
}
