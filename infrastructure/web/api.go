package web

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// OrderAPI type
type OrderAPI struct {
	Handler *OrderHandler
}

// Serve func
func (api OrderAPI) Serve() {
	router := mux.NewRouter()

	router.Methods("GET").Path("/").Handler(http.RedirectHandler("/swagger-ui/", http.StatusPermanentRedirect))
	router.PathPrefix("/api-docs/").Handler(http.StripPrefix("/api-docs/", http.FileServer(http.Dir("./infrastructure/web/api/"))))
	router.PathPrefix("/swagger-ui/").Handler(http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./infrastructure/web/swagger-ui/"))))

	router.Methods("POST").Path("/orders").HandlerFunc(api.Handler.CreateOrder)
	router.Methods("PUT").Path("/orders/{id}").HandlerFunc(api.Handler.UpdateOrder)

	log.Fatal(http.ListenAndServe(":8080", router))
}
