package domain

// ShippingOrderService type
type ShippingOrderService struct {
	Repository ShippingOrderRepository
}

// ShippingOrderRepository interface
type ShippingOrderRepository interface {
	save(order *ShippingOrder)
}

// CreateOrder func
func (orderService ShippingOrderService) CreateOrder(order *ShippingOrder) {
	// TODO validation
	orderService.Repository.save(order)
}
