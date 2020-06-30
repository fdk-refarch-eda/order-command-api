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

	orderEventProducer := event.NewTransactionalKafkaProducer(
		config.Topics.OrderEvents,
		config.Kafka.OrderEventProducerProperties,
		proto.MarshalOrderEvent,
	)
	defer orderEventProducer.Close()

	orderCommandConsumer := event.NewKafkaConsumer(
		config.Topics.OrderCommands,
		config.Kafka.OrderCommandConsumerProperties,
		&domain.OrderCommandProcessor{
			Repository:        postgresqlOrderRepo,
			OrderEventEmitter: orderEventProducer,
		},
		proto.UnmarshalOrderCommand,
	)

	defer orderCommandConsumer.Close()
	go orderCommandConsumer.Start()

	orderEventConsumer := event.NewKafkaConsumer(
		config.Topics.OrderEvents,
		config.Kafka.OrderEventConsumerProperties,
		&domain.OrderEventProcessor{},
		proto.UnmarshalOrderEvent,
	)

	defer orderEventConsumer.Close()
	go orderEventConsumer.Start()

	orderCommandProducer := event.NewKafkaProducer(
		config.Topics.OrderCommands,
		config.Kafka.OrderCommandProducerProperties,
		proto.MarshalOrderCommand,
	)
	defer orderCommandProducer.Close()

	api := &web.OrderAPI{
		Handler: &web.OrderHandler{
			Adapter: &adapter.OrderHandler{
				Service: &domain.ShippingOrderService{
					CommandEmitter: orderCommandProducer,
				},
			},
		},
	}

	api.Serve()
}
