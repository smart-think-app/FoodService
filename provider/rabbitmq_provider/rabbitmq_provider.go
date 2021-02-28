package rabbitmq_provider

import (
	"fmt"
	"github.com/streadway/amqp"
)

func ConnectRabbitMQ() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil
	}
	fmt.Println("Connect RabbitMQ Success")
	return conn
}

func QueueDeclare(ch *amqp.Channel) error {
	_, err := ch.QueueDeclare(
		"test", false, false, false, false, nil)
	if err != nil {
		return err
	}
	return nil
}
