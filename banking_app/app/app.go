package app

import (
	"log"
	"net/http"
	"github.com/gorilla/mux" // alternative: gin-gonic/gin
	"github.com/puttarajkoliwad/go_projects/banking_app/service"
	"github.com/puttarajkoliwad/go_projects/banking_app/domain"
	"github.com/jmoiron/sqlx"
	"time"
)

func Start() {
	// init multiplexer
	router := mux.NewRouter()

	// wiring
	dbClient := getDbClient()

	ch := &CustomerHandlers{
		// service.NewCustomerService(domain.NewCustomerRepositoryStub())
		service.NewCustomerService(domain.NewCustomerRepositoryDB(dbClient)),
	}

		// service.NewCustomerService(domain.NewAccountRepositoryDB(dbClient)),


	// define routes
	router.HandleFunc("/greetings", greetingsHandler).Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:\\d+}", ch.getCustomer).Methods(http.MethodGet)
	// router.HandleFunc("/customers/{customer_id:\\d+}", getCustomer).Methods(http.MethodGet)
	// router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	// listen and serve
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func getDbClient() *sqlx.DB {
	client, err := sqlx.Open("mysql", "root@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}