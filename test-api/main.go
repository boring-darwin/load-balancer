package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", welcome)
	http.HandleFunc("/health", isHealthy)

	http.ListenAndServe(":8083", nil)

}

func welcome(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to the 8083")
}

func isHealthy(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "I am healthy")
}
