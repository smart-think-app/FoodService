package provider

type IRabbitMQSupport interface {
	Publish(queueName string, body string) error
}
