package module_cmd

import (
	"FoodService/provider/postgres_provider"
	"FoodService/provider/rabbitmq_provider"
	"FoodService/router"
	"database/sql"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/streadway/amqp"
)

var (
	db       *sql.DB
	rabbitmq *amqp.Connection
)

func RunAPI() error {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	db = postgres_provider.ConnectPostgres()
	if db == nil {
		return errors.New("Connect Database Fail. ")
	}
	rabbitmq = rabbitmq_provider.ConnectRabbitMQ()
	if rabbitmq == nil {
		return errors.New("Connect RabbitMQ Fail. ")
	}
	ch, err := rabbitmq.Channel()
	if err != nil {
		return err
	}
	errQueueDeclare := rabbitmq_provider.QueueDeclare(ch)
	if errQueueDeclare != nil {
		return errQueueDeclare
	}
	router.FoodRouter(db, ch, e)
	router.DrinkRouter(e)
	// Start server
	e.Logger.Fatal(e.Start(":3001"))
	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Print(err.Error())
		}
		fmt.Print("End!")
	}()
	return nil
}
