package domain

import (
	"fmt"
	"log"
	"reflect"
)

// EventEmitter interface
type EventEmitter interface {
	Emit(event interface{})
}

// CreateOrderCommand model
type CreateOrderCommand struct {
	OrderID              string
	CustomerID           string
	ProductID            string
	Quantity             int
	ExpectedDeliveryDate string
	PickupDate           string
	PickupAddress        Address
	DestinationAddress   Address
	Status               string
}

// EventListener interface
type EventListener interface {
	Handle(event interface{})
}

// OrderCommandAgent type
type OrderCommandAgent struct {
	Repository ShippingOrderRepository
}

// Handle func
func (agent OrderCommandAgent) Handle(event interface{}) {
	switch event.(type) {
	case CreateOrderCommand:
		log.Println(fmt.Sprintf("Received CreateOrderCommand: %+v", event))
		agent.Repository.Save(mapToShippingOrder(event.(CreateOrderCommand)))
	default:
		log.Println(fmt.Sprintf("Received unknown event (%s). Ignoring...", reflect.TypeOf(event)))
	}
}

func mapToShippingOrder(command CreateOrderCommand) ShippingOrder {
	return ShippingOrder{
		OrderID:              command.OrderID,
		ProductID:            command.ProductID,
		CustomerID:           command.CustomerID,
		Quantity:             command.Quantity,
		PickupAddress:        command.PickupAddress,
		PickupDate:           command.PickupDate,
		DestinationAddress:   command.DestinationAddress,
		ExpectedDeliveryDate: command.ExpectedDeliveryDate,
	}
}
