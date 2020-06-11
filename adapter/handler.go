package adapter

import (
	"github.com/fdk-refarch-eda/order-service/order-command-service/domain"
)

// OrderHandler type
type OrderHandler struct {
	Service *domain.ShippingOrderService
}

// CreateOrder handler
func (orderHandler OrderHandler) CreateOrder(createOrderRequest CreateOrderRequest) CreateOrderResponse {
	order := toShippingOrder(createOrderRequest)
	orderHandler.Service.CreateOrder(&order)
	return toCreateOrderResponse(order)
}

// UpdateOrder handler
func (orderHandler OrderHandler) UpdateOrder(updateOrderRequest UpdateOrderRequest) {}
