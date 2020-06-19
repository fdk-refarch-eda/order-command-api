package infrastructure

import (
	"fmt"
	"log"
	"strings"

	"github.com/fdk-refarch-eda/order-service/order-command-service/infrastructure/database"
	"github.com/spf13/viper"
)

// BindPostgresqlConfig func
func BindPostgresqlConfig() *database.PostgresqlConfig {
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

	log.Println(fmt.Sprintf("Working with db config: %+v", *dbConfig))

	return dbConfig
}
