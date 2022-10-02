package redis

import (
	"context"
	"encoding/json"
)

type MessageQueue struct {
	Sender   string
	Receiver string
	Message  string
}

var ctx = context.Background()

func SetRedis(key, value string) {
	rdb := RedisConf()

	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func GetRedis(key string) []MessageQueue {
	rdb := RedisConf()
	data := rdb.Get(ctx, key).Val()
	data_json := []MessageQueue{}
	json.Unmarshal([]byte(data), &data_json)
	return data_json
}
