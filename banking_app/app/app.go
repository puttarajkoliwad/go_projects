package app

import (
	"log"
	"net/http"
)

func Start() {
	// init multiplexer
	mux := http.NewServeMux()

	// define routes
	mux.HandleFunc("/greetings", GreetingsHandler)
	mux.HandleFunc("/customers", GetAllCustomersHandler)

	// listen and serve
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}