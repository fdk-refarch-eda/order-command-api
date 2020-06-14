package main

import (
	"github.com/fdk-refarch-eda/order-service/order-command-service/adapter"
	"github.com/fdk-refarch-eda/order-service/order-command-service/domain"
	"github.com/fdk-refarch-eda/order-service/order-command-service/infrastructure/database"
	"github.com/fdk-refarch-eda/order-service/order-command-service/infrastructure/event"
	"github.com/fdk-refarch-eda/order-service/order-command-service/infrastructure/web"
)

func main() {
	commandListener := &event.SimpleEventBusListener{
		Processor: &domain.OrderCommandProcessor{
			Repository: &database.InMemoryShippingOrderRepository{},
		},
	}

	commandListener.Listen()

	api := &web.OrderAPI{
		Handler: &web.OrderHandler{
			Adapter: &adapter.OrderHandler{
				Service: &domain.ShippingOrderService{
					CommandEmitter: &event.SimpleEventBusEmitter{},
				},
			},
		},
	}

	api.Serve()
}
