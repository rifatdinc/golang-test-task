package client

import amqp "github.com/rabbitmq/amqp091-go"

func NewChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://user:password@localhost:7001/")
	FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	return ch
}
