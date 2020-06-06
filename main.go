package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Methods("POST").Path("/orders").HandlerFunc(CreateOrder)
	router.Methods("PUT").Path("/orders/{id}").HandlerFunc(UpdateOrder)

	log.Fatal(http.ListenAndServe(":8080", router))
}
