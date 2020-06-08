package main

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
	PickupAddress        Address
	PickupDate           string
	DestinationAddress   Address
	ExpectedDeliveryDate string
	Status               string
}
