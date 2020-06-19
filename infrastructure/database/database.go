package database

import (
	"fmt"
	"log"

	"github.com/fdk-refarch-eda/order-service/order-command-service/domain"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// InMemoryShippingOrderRepository type
type InMemoryShippingOrderRepository struct{}

var store map[string]domain.ShippingOrder = make(map[string]domain.ShippingOrder)

// Save func
func (repo InMemoryShippingOrderRepository) Save(order domain.ShippingOrder) error {
	store[order.OrderID] = order
	log.Println(fmt.Sprintf("Stored order: %+v", order))
	return nil
}

// PostgresqlShippingOrderRepository type
type postgresqlShippingOrderRepository struct {
	db *pg.DB
}

// PostgresqlConfig type
type PostgresqlConfig struct {
	Address  string
	Username string
	Password string
	Database string
}

// NewPostgresqlShippingOrderRepository func
func NewPostgresqlShippingOrderRepository(config *PostgresqlConfig) *postgresqlShippingOrderRepository {
	db := pg.Connect(&pg.Options{
		Addr:     config.Address,
		User:     config.Username,
		Password: config.Password,
		Database: config.Database,
	})

	err := createSchema(db)
	if err != nil {
		panic(err)
	}

	return &postgresqlShippingOrderRepository{db: db}
}

// createSchema creates database schema for User and Story models.
func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*domain.ShippingOrder)(nil),
	}

	for _, model := range models {
		err := db.CreateTable(model, &orm.CreateTableOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}

// Save func
func (repo postgresqlShippingOrderRepository) Save(order domain.ShippingOrder) {
	log.Println(fmt.Sprintf("Storing order: %+v", order))
	repo.db.Insert(&order)
}

// Close func
func (repo postgresqlShippingOrderRepository) Close() error {
	return repo.db.Close()
}
