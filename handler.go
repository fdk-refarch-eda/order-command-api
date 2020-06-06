package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateOrder Handler
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "CreateOrder Handler triggered...")
}

// UpdateOrder Handler
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Fprintln(w, fmt.Sprintf("UpdateOrder Handler for id=%v triggered...", id))
}
