package adapter

import (
	"fmt"

	"github.com/fdk-refarch-eda/order-service/order-command-service/domain"
	"github.com/go-playground/validator/v10"
)

// OrderHandler type
type OrderHandler struct {
	Service *domain.ShippingOrderService
}

var validate *validator.Validate = validator.New()

// CreateOrder handler
func (orderHandler OrderHandler) CreateOrder(createOrderRequest CreateOrderRequest) (CreateOrderResponse, *ErrorResponse) {
	if err := validate.Struct(createOrderRequest); err != nil {
		errors := []Error{}
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, Error{
				Message: fmt.Sprintf("Value %v for field %s is invalid.", e.Value(), e.Field()),
			})
		}
		return CreateOrderResponse{}, &ErrorResponse{Errors: errors}
	}

	command := mapToCreateOrderCommand(createOrderRequest)
	orderHandler.Service.CreateOrder(&command)
	return mapToCreateOrderResponse(command), nil
}

// UpdateOrder handler
func (orderHandler OrderHandler) UpdateOrder(updateOrderRequest UpdateOrderRequest) {}
