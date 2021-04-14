package rabbitmq_provider

import (
	"FoodService/consume/food_consume_service"
	"FoodService/interface/consume"
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

func QueueConsume(ch *amqp.Channel) error {

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
	forever := make(chan bool)
	for _, item := range queueConfigList {
		var consume food_consume_service.FoodConsume
		errProcess := processQueue(item.QueueName, ch, consume)
		if errProcess != nil {
			fmt.Print(errProcess.Error())
		}
	}

	<-forever
	return nil
}

func processQueue(queueName string, ch *amqp.Channel, consume consume.IBasicConsume) error {
	msgs, err := ch.Consume(
		queueName, "", false, false, false, false, nil)
	if err != nil {
		return err
	}
	go func() {
		for d := range msgs {
			fmt.Printf("Received a message from queue %s with message %s \n", queueName, d.Body)
			err := consume.Run(d.Body)
			if err != nil {
				fmt.Printf(err.Error())
			}
			errAck := d.Ack(true)
			if errAck != nil {
				fmt.Print(errAck.Error())
			}
		}
	}()
	return nil
}
