package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fdk-refarch-eda/order-service/order-command-service/infrastructure"
	"github.com/fdk-refarch-eda/order-service/order-command-service/infrastructure/events"
	"github.com/gorilla/mux"
)

var orderCommandProducer *infrastructure.OrderCommandProducer = initOrderCommandProducer()

func initOrderCommandProducer() *infrastructure.OrderCommandProducer {
	ocp, err := infrastructure.NewOrderCommandProducer()
	if err != nil {
		log.Fatal(err)
	}
	return ocp
}

// CreateOrder Handler
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateOrder Handler triggered...")
	orderCommandProducer.Emit(events.CreateOrderCommand{
		TimestampInMillis: time.Now().UnixNano() / int64(time.Millisecond),
		Payload: events.OrderEventPayload{
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
	orderCommandProducer.Emit(events.UpdateOrderCommand{
		TimestampInMillis: time.Now().UnixNano() / int64(time.Millisecond),
		Payload: events.OrderEventPayload{
			OrderID:    id,
			ProductID:  "123",
			CustomerID: "456",
		},
	})
}
