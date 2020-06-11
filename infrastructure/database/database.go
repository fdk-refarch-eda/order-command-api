package database

import (
	"github.com/fdk-refarch-eda/order-service/order-command-service/domain"
	"github.com/google/uuid"
)

// InMemoryShippingOrderRepository type
type InMemoryShippingOrderRepository struct {
	domain.ShippingOrderRepository
}

var store map[string]*domain.ShippingOrder = make(map[string]*domain.ShippingOrder)

func (repo InMemoryShippingOrderRepository) save(order *domain.ShippingOrder) {
	order.OrderID = uuid.New().String()
	store[order.OrderID] = order
}
