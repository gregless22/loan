package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":7777", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// return the handler object
	fmt.Fprintf(w, "Hi welcome to the servder")
}
