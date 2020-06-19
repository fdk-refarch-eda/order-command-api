package main

import (
	"github.com/fdk-refarch-eda/order-service/order-command-service/adapter"
	"github.com/fdk-refarch-eda/order-service/order-command-service/domain"
	"github.com/fdk-refarch-eda/order-service/order-command-service/infrastructure"
	"github.com/fdk-refarch-eda/order-service/order-command-service/infrastructure/database"
	"github.com/fdk-refarch-eda/order-service/order-command-service/infrastructure/event"
	"github.com/fdk-refarch-eda/order-service/order-command-service/infrastructure/web"
)

// Config type
type Config struct {
	PostgresqlConfig *database.PostgresqlConfig
}

func main() {
	const orderCommandsTopicName = "order-commands"
	const orderEventsTopicName = "orders"

	config := &Config{
		PostgresqlConfig: infrastructure.BindPostgresqlConfig(),
	}

	postgresqlOrderRepo := database.NewPostgresqlShippingOrderRepository(config.PostgresqlConfig)
	defer postgresqlOrderRepo.Close()

	commandListener := &event.SimpleEventBusListener{
		Topic: orderCommandsTopicName,
		Processor: &domain.OrderCommandProcessor{
			Repository: postgresqlOrderRepo,
			OrderEventEmitter: &event.SimpleEventBusEmitter{
				Topic: orderEventsTopicName,
			},
		},
	}

	commandListener.Listen()

	orderEventListener := &event.SimpleEventBusListener{
		Topic:     orderEventsTopicName,
		Processor: &domain.OrderEventProcessor{},
	}

	orderEventListener.Listen()

	api := &web.OrderAPI{
		Handler: &web.OrderHandler{
			Adapter: &adapter.OrderHandler{
				Service: &domain.ShippingOrderService{
					CommandEmitter: &event.SimpleEventBusEmitter{
						Topic: orderCommandsTopicName,
					},
				},
			},
		},
	}

	api.Serve()
}
