package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// TODO import the handlers
	http.HandleFunc("/", handler)
	http.HandleFunc("/create", create)
	http.HandleFunc("/read", read)

	log.Fatal(http.ListenAndServe(":7777", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// return the handler object
	fmt.Fprintf(w, "Hi welcome to the servder")
}

func create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this will eventually write")
}

// TODO Read
func read(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this will eventually read")
}

// TODO Update

// TODO Delete
