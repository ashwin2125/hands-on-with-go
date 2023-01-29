package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Your solution goes here. Good luck!
	router()
}

func router() {
	http.HandleFunc("/hello", serverFunction)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serverFunction(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello, "+name)
}
