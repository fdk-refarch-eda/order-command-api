package main

import (
	"strings"

	"github.com/fdk-refarch-eda/order-service/order-command-service/infrastructure/database"
	"github.com/spf13/viper"
)

// Config type
type config struct {
	TopicNames       Topics
	PostgresqlConfig database.PostgresqlConfig
}

// Topics type
type Topics struct {
	OrderCommands string
	OrderEvents   string
}

// NewConfig ctor
func NewConfig() config {
	return config{
		TopicNames:       bindTopicNames(),
		PostgresqlConfig: bindPostgresqlConfig(),
	}
}

func bindTopicNames() Topics {
	topics := &Topics{
		OrderCommands: "order-commands",
		OrderEvents:   "orders",
	}

	v := viper.New()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetEnvPrefix("topics")
	v.BindEnv("OrderCommands")
	v.BindEnv("OrderEvents")
	v.Unmarshal(topics)

	return *topics
}

func bindPostgresqlConfig() database.PostgresqlConfig {
	dbConfig := &database.PostgresqlConfig{
		Address:  ":5432",
		Username: "postgres",
		Database: "postgres",
	}

	v := viper.New()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetEnvPrefix("pg")
	v.BindEnv("address")
	v.BindEnv("username")
	v.BindEnv("password")
	v.BindEnv("database")
	v.Unmarshal(dbConfig)

	return *dbConfig
}
