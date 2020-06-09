package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// CreateOrderRequest model
// swagger:model
type CreateOrderRequest struct {
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
	var createOrderRequest CreateOrderRequest
	json.NewDecoder(r.Body).Decode(&createOrderRequest)
	log.Println(fmt.Sprintf("CreateOrder Handler triggered with (%+v) ...", createOrderRequest))

	orderCommandProducer.Emit(CreateOrderCommand{
		TimestampInMillis: time.Now().UnixNano() / int64(time.Millisecond),
		Payload:           toEventPayload(createOrderRequest),
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

func toEventPayload(createOrderRequest CreateOrderRequest) OrderEventPayload {
	return OrderEventPayload{
		OrderID:              uuid.New().String(),
		ProductID:            createOrderRequest.ProductID,
		CustomerID:           createOrderRequest.CustomerID,
		Quantity:             createOrderRequest.Quantity,
		PickupAddress:        createOrderRequest.PickupAddress,
		PickupDate:           createOrderRequest.PickupDate,
		DestinationAddress:   createOrderRequest.DestinationAddress,
		ExpectedDeliveryDate: createOrderRequest.ExpectedDeliveryDate,
		Status:               "to-be-created",
	}
}
