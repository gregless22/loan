package router

import (
	"fmt"
	"net/http"

	"github.com/gregless22/loan/middleware"
)

// Router will return a new servemux instance
func Router() *http.ServeMux {
	router := http.ServeMux{}

	router.HandleFunc("/loans/", loans)

	return &router
}

// this will parse the header and route to the appropriate function
func loans(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case "GET":
		middleware.ReadLoans(w, r)
	case "POST":
		middleware.CreateLoan(w, r)
	case "DELETE":
		middleware.DeleteLoan(w, r)
	case "PUT":
		middleware.UpdateLoan(w, r)
	}
}
