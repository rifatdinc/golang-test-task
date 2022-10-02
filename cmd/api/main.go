package main

import (
	v1 "twitch_chat_analysis/internal/controller/http/v1"
	consumer "twitch_chat_analysis/pkg/rabbitmq/client"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/message", v1.NewMessageHandler)
	r.GET("/message/list", v1.MessageList)
	go consumer.NewConsumerQueue()
	r.Run()
}
