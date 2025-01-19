package app

import (
	"fmt"
	"net/http"
	"encoding/json"
	"encoding/xml"
	"github.com/gorilla/mux"
)

type Customer struct {
	Name string `json:"name" xml:"name"`
	City string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode", xml:"zipcode"`
}

func greetingsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func getAllCustomers() []Customer {
	return []Customer{
		{"Rahul", "Blr", "567867"},
		{"Chetan", "blr", "567867"},
	}
}

func getAllCustomersHandler(w http.ResponseWriter, r *http.Request) {
	customers := getAllCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, vars["customer_id"])
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create customer")
}