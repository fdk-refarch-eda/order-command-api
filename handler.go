package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var orderCommandProducer *OrderCommandProducer = initOrderCommandProducer()

func initOrderCommandProducer() *OrderCommandProducer {
	ocp, err := NewOrderCommandProducer()
	if err != nil {
		log.Fatal(err)
	}
	return ocp
}

// CreateOrder Handler
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateOrder Handler triggered...")
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
