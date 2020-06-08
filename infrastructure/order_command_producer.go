package infrastructure

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/fdk-refarch-eda/order-service/order-command-service/infrastructure/events"
)

// OrderCommandProducer type
type OrderCommandProducer struct {
	producer *kafka.Producer
}

// NewOrderCommandProducer ctor
func NewOrderCommandProducer() (*OrderCommandProducer, error) {
	delegate, err := kafka.NewProducer(&kafka.ConfigMap{})

	if err != nil {
		return nil, err
	}

	p := &OrderCommandProducer{
		producer: delegate,
	}

	return p, nil
}

// Emit func
func (ocp OrderCommandProducer) Emit(command interface{}) {
	switch command.(type) {
	case events.CreateOrderCommand:
		log.Println(fmt.Sprintf("Emitting create-order-command (%+v)...", command))
	case events.UpdateOrderCommand:
		log.Println(fmt.Sprintf("Emitting update-order-command (%+v)...", command))
	default:
		log.Println("Received unknown command. Ignoring...")
	}
}
