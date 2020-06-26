package event

import (
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
	topic      string
	delegate   *kafka.Producer
	serializer func(event domain.Event) ([]byte, error)
}

// NewKafkaProducer ctor
func NewKafkaProducer(
	topic string,
	config map[string]interface{},
	serializer func(event domain.Event) ([]byte, error),
) *kafkaProducer {
	producer, err := kafka.NewProducer(toKafkaConfig(config))
	if err != nil {
		log.Fatal("Error while trying to create kafka producer", err)
	}

	return &kafkaProducer{
		topic:      topic,
		delegate:   producer,
		serializer: serializer,
	}
}

func (producer kafkaProducer) Emit(event domain.Event) {
	payload, err := producer.serializer(event)
	if err == nil {
		producer.delegate.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &producer.topic,
				Partition: kafka.PartitionAny,
			},
			Key:   []byte(event.ID()),
			Value: payload,
		}, nil)
	} else {
		log.Println("Error while trying to marshal event", event, err)
	}
}

func (producer kafkaProducer) Close() {
	producer.delegate.Close()
}

type kafkaConsumer struct {
	topic        string
	delegate     *kafka.Consumer
	processor    domain.EventProcessor
	deserializer func(data []byte) (domain.Event, error)
}

// NewKafkaConsumer ctor
func NewKafkaConsumer(
	topic string,
	config map[string]interface{},
	processor domain.EventProcessor,
	deserializer func(data []byte) (domain.Event, error),
) *kafkaConsumer {
	consumer, err := kafka.NewConsumer(toKafkaConfig(config))
	if err != nil {
		log.Fatal("Error while trying to create kafka consumer", err)
	}

	return &kafkaConsumer{
		topic:        topic,
		delegate:     consumer,
		processor:    processor,
		deserializer: deserializer,
	}
}

func (consumer kafkaConsumer) Start() {
	if err := consumer.delegate.Subscribe(consumer.topic, nil); err != nil {
		log.Fatal("Error while trying to subscribe to topic", consumer.topic, err)
	}

	for {
		msg, err := consumer.delegate.ReadMessage(-1)
		if err == nil {
			if event, err := consumer.deserializer(msg.Value); err == nil {
				consumer.processor.Process(event)
			} else {
				log.Println("Error while trying to unmarshal", msg.Value, err)
			}
		} else {
			log.Fatal("Error while trying to read messags from topic", consumer.topic, err)
			break
		}
	}

	consumer.delegate.Close()
}

func toKafkaConfig(config map[string]interface{}) *kafka.ConfigMap {
	c := kafka.ConfigMap{}

	for k, v := range config {
		c[k] = v
	}

	return &c
}
