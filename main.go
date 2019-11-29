package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/gql", handler)
	database.openJSON()
	log.Fatal(http.ListenAndServe(":7777", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HI welcome to the servder")
}
