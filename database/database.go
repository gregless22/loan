package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

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


