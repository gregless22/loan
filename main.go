package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gregless22/loan/database"
)

func main() {
	http.HandleFunc("/gql", handler)
	database.OpenJSON()
	log.Fatal(http.ListenAndServe(":7777", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HI welcome to the servder")
}
