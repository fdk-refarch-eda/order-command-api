package adapter

import "github.com/fdk-refarch-eda/order-service/order-command-service/domain"

// CreateOrderRequest model
// swagger:model
type CreateOrderRequest struct {
	CustomerID           string  `json:"customer_id"`
	ProductID            string  `json:"product_id"`
	Quantity             int     `json:"quantity"`
	ExpectedDeliveryDate string  `json:"expected_delivery_date"`
	PickupDate           string  `json:"pickup_date"`
	PickupAddress        Address `json:"pickup_address"`
	DestinationAddress   Address `json:"destination_address"`
}

// Address model
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}

// CreateOrderResponse model
// swagger:model
type CreateOrderResponse struct {
	OrderID              string  `json:"order_id"`
	CustomerID           string  `json:"customer_id"`
	ProductID            string  `json:"product_id"`
	Quantity             int     `json:"quantity"`
	ExpectedDeliveryDate string  `json:"expected_delivery_date"`
	PickupDate           string  `json:"pickup_date"`
	PickupAddress        Address `json:"pickup_address"`
	DestinationAddress   Address `json:"destination_address"`
}

// UpdateOrderRequest model
// swagger:model
type UpdateOrderRequest struct {
	OrderID              string  `json:"order_id"`
	CustomerID           string  `json:"customer_id"`
	ProductID            string  `json:"product_id"`
	Quantity             int     `json:"quantity"`
	ExpectedDeliveryDate string  `json:"expected_delivery_date"`
	PickupDate           string  `json:"pickup_date"`
	PickupAddress        Address `json:"pickup_address"`
	DestinationAddress   Address `json:"destination_address"`
}

func toShippingOrder(createOrderRequest CreateOrderRequest) domain.ShippingOrder {
	return domain.ShippingOrder{
		ProductID:            createOrderRequest.ProductID,
		CustomerID:           createOrderRequest.CustomerID,
		Quantity:             createOrderRequest.Quantity,
		PickupAddress:        toDomainAddress(createOrderRequest.PickupAddress),
		PickupDate:           createOrderRequest.PickupDate,
		DestinationAddress:   toDomainAddress(createOrderRequest.DestinationAddress),
		ExpectedDeliveryDate: createOrderRequest.ExpectedDeliveryDate,
		Status:               "to-be-created",
	}
}

func toDomainAddress(address Address) domain.Address {
	return domain.Address{
		Street:  address.Street,
		City:    address.City,
		Country: address.Country,
		State:   address.State,
		ZipCode: address.ZipCode,
	}
}

func toAddress(address domain.Address) Address {
	return Address{
		Street:  address.Street,
		City:    address.City,
		Country: address.Country,
		State:   address.State,
		ZipCode: address.ZipCode,
	}
}

func toCreateOrderResponse(order domain.ShippingOrder) CreateOrderResponse {
	return CreateOrderResponse{
		OrderID:              order.OrderID,
		CustomerID:           order.CustomerID,
		ProductID:            order.ProductID,
		Quantity:             order.Quantity,
		ExpectedDeliveryDate: order.ExpectedDeliveryDate,
		PickupDate:           order.PickupDate,
		PickupAddress:        toAddress(order.PickupAddress),
		DestinationAddress:   toAddress(order.DestinationAddress),
	}
}
