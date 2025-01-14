package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"encoding/xml"
)

type Customer struct {
	Name string `json:"name" xml:"name"`
	City	string	`json:"city" xml:"city"`
	Zipcode string	`json:"zipcode" xml:"zipcode"`
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Ashish", "New Delhi", "110075"},
		{"Rahul", "New Delhi", "110075"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func main() {
	// routes
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/customers", getAllCustomers)

	// listen and serve
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}