package v1

import (
	"errors"

	mq_client "twitch_chat_analysis/pkg/rabbitmq/client"
	"twitch_chat_analysis/pkg/redis"

	"github.com/gin-gonic/gin"
)

func NewMessageHandler(c *gin.Context) {
	jsons := make([]byte, c.Request.ContentLength)
	if _, err := c.Request.Body.Read(jsons); err != nil {
		if err.Error() != "EOF" {
			c.AbortWithError(400, errors.New("failed to declare a queue"))
		}
	}
	err := mq_client.NewSenderQueue(jsons)

	if err != nil {
		c.AbortWithError(400, errors.New("failed to declare a queue"))
	} else {
		c.Status(200)
	}
}

func MessageList(c *gin.Context) {
	params := c.Request.URL.Query()
	sender_key := ""
	for _, v := range params {
		sender_key = string(v[0])
	}

	c.JSON(200, redis.GetRedis(sender_key))

}
