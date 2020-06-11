package domain

// ShippingOrderService type
type ShippingOrderService struct {
	Repository ShippingOrderRepository
}

// ShippingOrderRepository interface
type ShippingOrderRepository interface {
	Save(order *ShippingOrder)
}

// CreateOrder func
func (orderService ShippingOrderService) CreateOrder(order *ShippingOrder) {
	// TODO validation
	orderService.Repository.Save(order)
}
