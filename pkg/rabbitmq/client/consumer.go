package client

import (
	"encoding/json"
	"fmt"
	"log"
	"twitch_chat_analysis/pkg/redis"
)

type MessageQueue struct {
	Sender   string
	Receiver string
	Message  string
}

type MessageList struct {
	Item []MessageQueue
}

func NewConsumerQueue() {
	ch := NewChannel()
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"message", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			jsonTostring := string(d.Body[:])
			data := MessageQueue{}
			json.Unmarshal([]byte(jsonTostring),&data)
			toString := string(d.Body)

			stringArray := []string{toString}
			justString := fmt.Sprint(stringArray)
			redis.SetRedis(data.Sender, justString)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever
}
