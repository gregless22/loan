package main

import (
	"log"
	"net/http"

	"github.com/gregless22/loan/server/models"
	"github.com/gregless22/loan/server/router"
)

func main() {
	// initialise the models in the database
	var l models.Loan
	l.Init()

	// add the router
	router := router.Router()

	log.Fatal(http.ListenAndServe(":3030", router))
}
