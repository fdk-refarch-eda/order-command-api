package events

import "github.com/fdk-refarch-eda/order-service/order-command-service/domain"

// CreateOrderCommand event
type CreateOrderCommand struct {
	TimestampInMillis int64
	Payload           OrderEventPayload
}

// UpdateOrderCommand event
type UpdateOrderCommand struct {
	TimestampInMillis int64
	Payload           OrderEventPayload
}

// OrderEventPayload type
type OrderEventPayload struct {
	OrderID              string
	ProductID            string
	CustomerID           string
	Quantity             int
	PickupAddress        domain.Address
	PickupDate           string
	DestinationAddress   domain.Address
	ExpectedDeliveryDate string
	Status               string
}
