package domain

// ShippingOrder entity
type ShippingOrder struct {
	OrderID              string
	ProductID            string
	CustomerID           string
	Quantity             int
	PickupAddress        Address
	PickupDate           string
	DestinationAddress   Address
	ExpectedDeliveryDate string
	Status               string
	VoyageID             string
	ContainerID          string
}

// Address model
type Address struct {
	Street  string
	City    string
	Country string
	State   string
	ZipCode string
}
