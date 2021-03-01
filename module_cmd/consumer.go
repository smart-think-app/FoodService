package module_cmd

import (
	"FoodService/provider/postgres_provider"
	"FoodService/provider/rabbitmq_provider"
	"errors"
	"fmt"
)

func RunConsumer() error {
	fmt.Print("Run Consumer")
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
	errQueueConsume := rabbitmq_provider.QueueConsume(ch)
	if errQueueConsume != nil {
		return errQueueConsume
	}
	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Print(err.Error())
		}
		errCh := ch.Close()
		if errCh != nil {
			fmt.Print(errCh.Error())
		}
		fmt.Print("End!")
	}()
	return nil
}
