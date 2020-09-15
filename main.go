package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gregless22/loan/models"
)

func main() {
	// initialise the models in the database
	var l models.Loan
	l.Init()

	http.HandleFunc("/", test)
	http.HandleFunc("/loans", loans)

	log.Fatal(http.ListenAndServe(":3030", nil))
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "made it to the test server")
}

func loans(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// var loan models.Loan

		// loans, err := loan.GetAll()

		// if err != nil {
		// 	fmt.Fprintf(w, "Error getting data: %s", err)
		// }

		// loansJSON, err := json.Marshal(loans)

		// if err != nil {
		// 	fmt.Fprintf(w, "Error decoding data: %s", err)
		// }

		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// w.Write(loansJSON)

	case "POST":
		// loan := models.Loan{}

		// // decode the incoming body
		// err := json.NewDecoder(r.Body).Decode(&loan)
		// if err != nil {
		// 	fmt.Fprintf(w, "Error decoding data: %s", err)
		// }

		// loan.Create()
		// // TODO write the loan to the database

		// // Send the reponse back that the loan has been added
		// loanJSON, err := json.Marshal(loan)
		// if err != nil {
		// 	fmt.Fprintf(w, "Error marshalling back to JSON")
		// }

		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// w.Write(loanJSON)
	case "DELETE":
		fmt.Fprintf(w, "TODO delete the loan")
	case "PUT":
		fmt.Fprintf(w, "TODO update the loan")
	}
}
