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

// EventProcessor interface
type EventProcessor interface {
	Process(event interface{})
}

// OrderCommandProcessor type
type OrderCommandProcessor struct {
	Repository ShippingOrderRepository
}

// Process func
func (orderCommandProcessor OrderCommandProcessor) Process(event interface{}) {
	switch event.(type) {
	case CreateOrderCommand:
		log.Println(fmt.Sprintf("Received CreateOrderCommand: %+v", event))
		shippingOrder := mapToShippingOrder(event.(CreateOrderCommand))
		orderCommandProcessor.Repository.Save(shippingOrder)
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
