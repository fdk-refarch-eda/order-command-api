package domain

import "github.com/google/uuid"

// ShippingOrderService type
type ShippingOrderService struct {
	CommandEmitter EventEmitter
}

// ShippingOrderRepository interface
type ShippingOrderRepository interface {
	Save(order ShippingOrder)
}

// CreateOrder func
func (orderService ShippingOrderService) CreateOrder(command *CreateOrderCommand) {
	command.OrderID = uuid.New().String()
	orderService.CommandEmitter.Emit(*command)
}
