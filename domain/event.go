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
}

// OrderCreatedEvent model
type OrderCreatedEvent struct {
	OrderID              string
	CustomerID           string
	ProductID            string
	Quantity             int
	ExpectedDeliveryDate string
	PickupDate           string
	PickupAddress        Address
	DestinationAddress   Address
}

func newOrderCreatedEvent(order ShippingOrder) OrderCreatedEvent {
	return OrderCreatedEvent{
		OrderID:              order.OrderID,
		ProductID:            order.ProductID,
		CustomerID:           order.CustomerID,
		Quantity:             order.Quantity,
		PickupAddress:        order.PickupAddress,
		PickupDate:           order.PickupDate,
		DestinationAddress:   order.DestinationAddress,
		ExpectedDeliveryDate: order.ExpectedDeliveryDate,
	}
}

// EventProcessor interface
type EventProcessor interface {
	Process(event interface{})
}

// OrderCommandProcessor type
type OrderCommandProcessor struct {
	Repository        ShippingOrderRepository
	OrderEventEmitter EventEmitter
}

// Process func
func (orderCommandProcessor OrderCommandProcessor) Process(event interface{}) {
	switch event.(type) {
	case CreateOrderCommand:
		log.Println(fmt.Sprintf("Received CreateOrderCommand: %+v", event))
		createOrderCommand := event.(CreateOrderCommand)
		shippingOrder := mapToShippingOrder(createOrderCommand)
		orderCommandProcessor.Repository.Save(shippingOrder)
		orderCreatedEvent := newOrderCreatedEvent(shippingOrder)
		orderCommandProcessor.OrderEventEmitter.Emit(orderCreatedEvent)
	default:
		log.Println(fmt.Sprintf("Received unknown event (%s). Ignoring...", reflect.TypeOf(event)))
	}
}

// OrderEventProcessor type
type OrderEventProcessor struct{}

// Process func
func (orderEventProcessor OrderEventProcessor) Process(event interface{}) {
	switch event.(type) {
	case OrderCreatedEvent:
		log.Println(fmt.Sprintf("Received OrderCreatedEvent: %+v", event))
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
