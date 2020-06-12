package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fdk-refarch-eda/order-service/order-command-service/adapter"
	"github.com/gorilla/mux"
)

// OrderHandler type
type OrderHandler struct {
	Adapter *adapter.OrderHandler
}

// CreateOrder Handler
func (orderHandler OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var createOrderRequest adapter.CreateOrderRequest
	json.NewDecoder(r.Body).Decode(&createOrderRequest)

	createOrderResponse, err := orderHandler.Adapter.CreateOrder(createOrderRequest)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(createOrderResponse)
	}
}

// UpdateOrder Handler
func (orderHandler OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	log.Println(fmt.Sprintf("UpdateOrder Handler for id=%v triggered...", id))

	var updateOrderRequest adapter.UpdateOrderRequest
	json.NewDecoder(r.Body).Decode(&updateOrderRequest)

	orderHandler.Adapter.UpdateOrder(updateOrderRequest)
}
