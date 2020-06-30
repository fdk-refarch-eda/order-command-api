package event

import (
	"context"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/fdk-refarch-eda/order-service/order-command-service/domain"
)

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

func (producer kafkaProducer) Emit(ctx context.Context, event domain.Event) {
	payload, err := producer.serializer(event)
	if err == nil {
		deliveryChan := make(chan kafka.Event)

		producer.delegate.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &producer.topic,
				Partition: kafka.PartitionAny,
			},
			Key:   []byte(event.ID()),
			Value: payload,
		}, deliveryChan)

		producedEvent := <-deliveryChan

		if err := producedEvent.(*kafka.Message).TopicPartition.Error; err != nil {
			log.Println("Error while trying to produce message with payload", payload, err)
		}

		close(deliveryChan)
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

	consumerMetaData, _ := consumer.delegate.GetConsumerGroupMetadata()

	for {
		msg, err := consumer.delegate.ReadMessage(-1)
		if err == nil {
			if event, err := consumer.deserializer(msg.Value); err == nil {

				tp := msg.TopicPartition
				tp.Offset = tp.Offset + 1

				ctx := context.WithValue(
					context.Background(),
					transactionMetaDataKey, transactionMetaData{
						consumerGroupMetaData: consumerMetaData,
						offsets:               []kafka.TopicPartition{tp},
					},
				)

				consumer.processor.Process(ctx, event)
			} else {
				log.Println("Error while trying to unmarshal", msg.Value, err)
			}
		} else {
			log.Fatal("Error while trying to read messags from topic ", consumer.topic, err)
			continue
		}
	}
}

func (consumer kafkaConsumer) Close() error {
	return consumer.delegate.Close()
}

func toKafkaConfig(config map[string]interface{}) *kafka.ConfigMap {
	c := kafka.ConfigMap{}

	for k, v := range config {
		c[k] = v
	}

	return &c
}

type transactionalKafkaProducer struct {
	topic      string
	delegate   *kafka.Producer
	serializer func(event domain.Event) ([]byte, error)
}

type key string

const transactionMetaDataKey key = "transaction-metadata"

type transactionMetaData struct {
	consumerGroupMetaData *kafka.ConsumerGroupMetadata
	offsets               []kafka.TopicPartition
}

// NewTransactionalKafkaProducer ctor
func NewTransactionalKafkaProducer(
	topic string,
	config map[string]interface{},
	serializer func(event domain.Event) ([]byte, error),
) *transactionalKafkaProducer {
	producer, err := kafka.NewProducer(toKafkaConfig(config))
	if err != nil {
		log.Fatal("Error while trying to create kafka producer", err)
	}

	producer.InitTransactions(nil)

	return &transactionalKafkaProducer{
		topic:      topic,
		delegate:   producer,
		serializer: serializer,
	}
}

func (producer transactionalKafkaProducer) Emit(ctx context.Context, event domain.Event) {
	payload, err := producer.serializer(event)
	if err == nil {
		producer.delegate.BeginTransaction()

		deliveryChan := make(chan kafka.Event)

		producer.delegate.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &producer.topic,
				Partition: kafka.PartitionAny,
			},
			Key:   []byte(event.ID()),
			Value: payload,
		}, deliveryChan)

		producedEvent := <-deliveryChan

		if err := producedEvent.(*kafka.Message).TopicPartition.Error; err == nil {
			metadata := ctx.Value(transactionMetaDataKey).(transactionMetaData)
			if err := producer.delegate.SendOffsetsToTransaction(ctx, metadata.offsets, metadata.consumerGroupMetaData); err == nil {
				producer.delegate.CommitTransaction(ctx)
			} else {
				log.Println("Error while sending offset to transaction", err)
				producer.delegate.AbortTransaction(ctx)
			}
		} else {
			log.Println("Error while trying to produce message with payload", payload, err)
			producer.delegate.AbortTransaction(ctx)
		}

		close(deliveryChan)
	} else {
		log.Println("Error while trying to marshal event", event, err)
	}
}

func (producer transactionalKafkaProducer) Close() {
	producer.delegate.Close()
}
