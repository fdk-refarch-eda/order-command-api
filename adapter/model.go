package adapter

import (
	"github.com/fdk-refarch-eda/order-service/order-command-service/domain"
)

// CreateOrderRequest model
// swagger:model
type CreateOrderRequest struct {
	CustomerID           string  `json:"customer_id" validate:"required"`
	ProductID            string  `json:"product_id" validate:"required"`
	Quantity             int     `json:"quantity" validate:"required,gt=0"`
	ExpectedDeliveryDate string  `json:"expected_delivery_date" validate:"required,datetime=2006-01-02T15:04Z"`
	PickupDate           string  `json:"pickup_date" validate:"required,datetime=2006-01-02T15:04Z"`
	PickupAddress        Address `json:"pickup_address" validate:"required"`
	DestinationAddress   Address `json:"destination_address" validate:"required"`
}

// Address model
type Address struct {
	Street  string `json:"street" validate:"required"`
	City    string `json:"city" validate:"required"`
	Country string `json:"country" validate:"required"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code" validate:"required"`
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
	OrderID              string  `json:"order_id" validate:"required"`
	CustomerID           string  `json:"customer_id" validate:"required"`
	ProductID            string  `json:"product_id" validate:"required"`
	Quantity             int     `json:"quantity" validate:"required,gt=0"`
	ExpectedDeliveryDate string  `json:"expected_delivery_date" validate:"required,datetime=2006-01-02T15:04Z"`
	PickupDate           string  `json:"pickup_date" validate:"required,datetime=2006-01-02T15:04Z"`
	PickupAddress        Address `json:"pickup_address" validate:"required"`
	DestinationAddress   Address `json:"destination_address" validate:"required"`
}

// ErrorResponse model
// swagger:model
type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

// Error model
// swagger:model
type Error struct {
	Message string `json:"message"`
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
