package event

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/asaskevich/EventBus"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/fdk-refarch-eda/order-service/order-command-service/domain"
)

var eventBus EventBus.Bus = EventBus.New()

// SimpleEventBusEmitter type
type SimpleEventBusEmitter struct {
	Topic string
}

// Emit Func
func (emitter SimpleEventBusEmitter) Emit(event domain.Event) {
	log.Println(fmt.Sprintf("Emitting %s: %+v to topic: %s", reflect.TypeOf(event), event, emitter.Topic))
	go eventBus.Publish(emitter.Topic, event)
}

// SimpleEventBusListener type
type SimpleEventBusListener struct {
	Topic     string
	Processor domain.EventProcessor
}

// Listen func
func (listener SimpleEventBusListener) Listen() {
	eventBus.Subscribe(listener.Topic, listener.Processor.Process)
}

type kafkaProducer struct {
	topic    string
	delegate *kafka.Producer
}

// NewKafkaProducer ctor
func NewKafkaProducer(topic string, config map[string]interface{}) *kafkaProducer {
	producer, err := kafka.NewProducer(toKafkaConfig(config))
	if err != nil {
		log.Fatal("Error while trying to create kafka producer", err)
	}

	return &kafkaProducer{
		topic:    topic,
		delegate: producer,
	}
}

func (producer kafkaProducer) Emit(event domain.Event) {
	payload, err := json.Marshal(event)
	if err == nil {
		producer.delegate.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &producer.topic,
				Partition: kafka.PartitionAny,
			},
			Key:   []byte(event.ID()),
			Value: payload,
		}, nil)
	}
}

func (producer kafkaProducer) Close() {
	producer.delegate.Close()
}

func toKafkaConfig(config map[string]interface{}) *kafka.ConfigMap {
	c := kafka.ConfigMap{}

	for k, v := range config {
		c[k] = v
	}

	return &c
}
