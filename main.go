package main

import (
	"github.com/fdk-refarch-eda/order-service/order-command-service/adapter"
	"github.com/fdk-refarch-eda/order-service/order-command-service/domain"
	"github.com/fdk-refarch-eda/order-service/order-command-service/infrastructure/database"
	"github.com/fdk-refarch-eda/order-service/order-command-service/infrastructure/web"
)

func main() {
	api := &web.OrderAPI{
		Handler: &web.OrderHandler{
			Adapter: &adapter.OrderHandler{
				Service: &domain.ShippingOrderService{
					Repository: &database.InMemoryShippingOrderRepository{},
				},
			},
		},
	}

	api.Serve()
}
