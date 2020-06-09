package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Methods("GET").Path("/").Handler(http.RedirectHandler("/swagger-ui/", http.StatusPermanentRedirect))
	router.PathPrefix("/api-docs/").Handler(http.StripPrefix("/api-docs/", http.FileServer(http.Dir("./api/"))))
	router.PathPrefix("/swagger-ui/").Handler(http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./web/swagger-ui/"))))

	router.Methods("POST").Path("/orders").HandlerFunc(CreateOrder)
	router.Methods("PUT").Path("/orders/{id}").HandlerFunc(UpdateOrder)

	log.Fatal(http.ListenAndServe(":8080", router))
}
