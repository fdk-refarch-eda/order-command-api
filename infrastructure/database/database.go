package database

import (
	"fmt"
	"log"

	"github.com/fdk-refarch-eda/order-service/order-command-service/domain"
)

// InMemoryShippingOrderRepository type
type InMemoryShippingOrderRepository struct{}

var store map[string]domain.ShippingOrder = make(map[string]domain.ShippingOrder)

// Save func
func (repo InMemoryShippingOrderRepository) Save(order domain.ShippingOrder) {
	store[order.OrderID] = order
	log.Println(fmt.Sprintf("Stored order: %+v", order))
}
