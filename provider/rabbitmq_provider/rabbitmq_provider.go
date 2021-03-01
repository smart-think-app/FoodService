package rabbitmq_provider

import (
	"FoodService/model/common_model"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"io/ioutil"
	"os"
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
	queueConfigList := make([]common_model.QueueConfigDto, 0)
	queueConfig, errConfig := os.Open("rabbitmq_config.json")
	if errConfig != nil {
		return errConfig
	}
	byteValue, _ := ioutil.ReadAll(queueConfig)
	errMarshal := json.Unmarshal(byteValue, &queueConfigList)
	if errMarshal != nil {
		return errMarshal
	}
	for _, item := range queueConfigList {
		if item.Status {
			_, err := ch.QueueDeclare(
				item.QueueName, false, false, false, false, nil)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
