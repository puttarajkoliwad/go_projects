package app

import (
	"fmt"
	"net/http"
	"encoding/json"
	"encoding/xml"
	"github.com/gorilla/mux"
	"github.com/puttarajkoliwad/go_projects/banking_app/service"
)

type Customer struct {
	Name string `json:"name" xml:"name"`
	City string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode", xml:"zipcode"`
}

type CustomerHandlers struct {
	svc service.CustomerService
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

func (ch CustomerHandlers) getAllCustomersHandler(w http.ResponseWriter, r *http.Request) {
	// customers := getAllCustomers()
	customers, _ := ch.svc.GetAllCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (ch CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// fmt.Fprintf(w, vars["customer_id"])

	customer, err := ch.svc.GetCustomer(vars["customer_id"])
	if err != nil {
		writeJsonResponse(w, err.Code, err.AsMessage())
		return
	}

	writeJsonResponse(w, http.StatusOK, customer)
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create customer")
}

func writeJsonResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}