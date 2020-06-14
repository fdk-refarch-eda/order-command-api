package domain

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
