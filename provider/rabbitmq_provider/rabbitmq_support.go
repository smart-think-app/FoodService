package rabbitmq_provider

import "github.com/streadway/amqp"

type RabbitMQSupport struct {
	Ch *amqp.Channel
}

func (rb RabbitMQSupport) Publish(queueName string, body string) error {
	err := rb.Ch.Publish("", queueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	if err != nil {
		return err
	}

	return nil
}
