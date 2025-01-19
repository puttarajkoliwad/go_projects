package app

import (
	"log"
	"net/http"
	"github.com/gorilla/mux" // alternative: gin-gonic/gin
)

func Start() {
	// init multiplexer
	router := mux.NewRouter()

	// wiring
	ch := &CustomerHandlers{
		service.NewCustomerService(domain.NewCustomerRepositoryStub())
	}

	// define routes
	router.HandleFunc("/greetings", greetingsHandler).Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.getAllCustomersHandler).Methods(http.MethodGet)
	// router.HandleFunc("/customers/{customer_id:\\d+}", getCustomer).Methods(http.MethodGet)
	// router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	// listen and serve
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
