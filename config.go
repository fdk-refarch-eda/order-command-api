package main

import (
	"log"

	"github.com/fdk-refarch-eda/order-service/order-command-service/infrastructure/database"
	"github.com/spf13/viper"
)

// Config type
type config struct {
	Topics           Topics
	PostgresqlConfig database.PostgresqlConfig
	Kafka            KafkaConfig
}

// Topics type
type Topics struct {
	OrderCommands string
	OrderEvents   string
}

// KafkaConfig type
type KafkaConfig struct {
	OrderCommandProducerProperties KafkaProperties
	OrderCommandConsumerProperties KafkaProperties
	OrderEventProducerProperties   KafkaProperties
	OrderEventConsumerProperties   KafkaProperties
}

// KafkaProperties type
type KafkaProperties map[string]interface{}

// NewConfig ctor
func NewConfig() config {
	c := config{}

	v := viper.NewWithOptions(viper.KeyDelimiter("#"))
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.BindEnv("PostgresqlConfig#Password", "PG_PASSWORD")

	if err := v.ReadInConfig(); err != nil {
		log.Fatal("Unable to read config file.", err)
	}

	if err := v.Unmarshal(&c); err != nil {
		log.Fatal("Unable to unmarshal config.", err)
	}

	return c
}
