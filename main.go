package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Starting Server ...")
	http.ListenAndServe(":8080", nil)
}

// Index page
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to index page")
}

// About page
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to About page")
}
