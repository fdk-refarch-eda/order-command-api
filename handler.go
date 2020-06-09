package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// OrderRequest model
// swagger:model
type OrderRequest struct {
	CustomerID           string
	ProductID            string
	Quantity             int
	ExpectedDeliveryDate string
	PickupDate           string
	PickupAddress        Address
	DestinationAddress   Address
}

var orderCommandProducer *OrderCommandProducer = NewOrderCommandProducer()

// CreateOrder Handler
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var orderRequest OrderRequest
	json.NewDecoder(r.Body).Decode(&orderRequest)
	log.Println(fmt.Sprintf("CreateOrder Handler triggered with (%+v) ...", orderRequest))
	orderCommandProducer.Emit(CreateOrderCommand{
		TimestampInMillis: time.Now().UnixNano() / int64(time.Millisecond),
		Payload: OrderEventPayload{
			OrderID:    "1",
			ProductID:  "123",
			CustomerID: "456",
		},
	})
}

// UpdateOrder Handler
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	log.Println(fmt.Sprintf("UpdateOrder Handler for id=%v triggered...", id))
	orderCommandProducer.Emit(UpdateOrderCommand{
		TimestampInMillis: time.Now().UnixNano() / int64(time.Millisecond),
		Payload: OrderEventPayload{
			OrderID:    id,
			ProductID:  "123",
			CustomerID: "456",
		},
	})
}
