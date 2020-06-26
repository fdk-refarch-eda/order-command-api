package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/fdk-refarch-eda/order-service/order-command-service/adapter"
	"github.com/fdk-refarch-eda/order-service/order-command-service/adapter/proto"
	"github.com/fdk-refarch-eda/order-service/order-command-service/domain"
	"github.com/fdk-refarch-eda/order-service/order-command-service/infrastructure/database"
	"github.com/fdk-refarch-eda/order-service/order-command-service/infrastructure/event"
	"github.com/fdk-refarch-eda/order-service/order-command-service/infrastructure/web"
)

func main() {
	config := NewConfig()
	log.Println("Working with config:", spew.Sdump(config))

	postgresqlOrderRepo := database.NewPostgresqlShippingOrderRepository(config.PostgresqlConfig)
	defer postgresqlOrderRepo.Close()

	commandConsumer := event.NewKafkaConsumer(
		config.Topics.OrderCommands,
		config.Kafka.OrderCommandConsumerProperties,
		&domain.OrderCommandProcessor{
			Repository: postgresqlOrderRepo,
			OrderEventEmitter: &event.SimpleEventBusEmitter{
				Topic: config.Topics.OrderEvents,
			},
		},
		proto.UnmarshalOrderCommand,
	)

	go commandConsumer.Start()

	orderEventListener := &event.SimpleEventBusListener{
		Topic:     config.Topics.OrderEvents,
		Processor: &domain.OrderEventProcessor{},
	}

	orderEventListener.Listen()

	kafkaProducer := event.NewKafkaProducer(
		config.Topics.OrderCommands,
		config.Kafka.OrderCommandProducerProperties,
		proto.MarshalOrderCommand,
	)
	defer kafkaProducer.Close()

	api := &web.OrderAPI{
		Handler: &web.OrderHandler{
			Adapter: &adapter.OrderHandler{
				Service: &domain.ShippingOrderService{
					CommandEmitter: kafkaProducer,
				},
			},
		},
	}

	api.Serve()
}
