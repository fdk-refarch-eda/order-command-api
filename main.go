package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/fdk-refarch-eda/order-service/order-command-service/adapter"
	"github.com/fdk-refarch-eda/order-service/order-command-service/domain"
	"github.com/fdk-refarch-eda/order-service/order-command-service/infrastructure/database"
	"github.com/fdk-refarch-eda/order-service/order-command-service/infrastructure/event"
	"github.com/fdk-refarch-eda/order-service/order-command-service/infrastructure/web"
)

func main() {
	config := NewConfig()
	log.Println(spew.Sprintf("Working with config: %+v", config))

	postgresqlOrderRepo := database.NewPostgresqlShippingOrderRepository(config.PostgresqlConfig)
	defer postgresqlOrderRepo.Close()

	commandListener := &event.SimpleEventBusListener{
		Topic: config.TopicNames.OrderCommands,
		Processor: &domain.OrderCommandProcessor{
			Repository: postgresqlOrderRepo,
			OrderEventEmitter: &event.SimpleEventBusEmitter{
				Topic: config.TopicNames.OrderEvents,
			},
		},
	}

	commandListener.Listen()

	orderEventListener := &event.SimpleEventBusListener{
		Topic:     config.TopicNames.OrderEvents,
		Processor: &domain.OrderEventProcessor{},
	}

	orderEventListener.Listen()

	api := &web.OrderAPI{
		Handler: &web.OrderHandler{
			Adapter: &adapter.OrderHandler{
				Service: &domain.ShippingOrderService{
					CommandEmitter: &event.SimpleEventBusEmitter{
						Topic: config.TopicNames.OrderCommands,
					},
				},
			},
		},
	}

	api.Serve()
}
