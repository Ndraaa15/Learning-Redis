package stream

import (
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   1,
})

var ctx = context.Background()

func TestPublishStream(t *testing.T){
	for i:= 0; i < 10; i++ {
		err := client.XAdd(ctx, &redis.XAddArgs{
			Stream: "members",
			Values: map[string]interface{}{
				"name": "Indra",
				"address": "Indonesia",
			},
		}).Err()

		assert.Nil(t, err)
	}
}

func TestCreateConsumerGroup(t *testing.T){
	client.XGroupCreate(ctx, "members", "consumer-group", "0").Err()
	client.XGroupCreateConsumer(ctx, "members", "consumer-group", "consumer-1")
	client.XGroupCreateConsumer(ctx, "members", "consumer-group", "consumer-1")
}

func TestReadStream(t *testing.T){
	streams, err := client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    "consumer-group",
		Consumer: "consumer-1",
		Streams:  []string{"members", ">"},
		Count:    10,
		Block:    5 * time.Second,
	}).Result()

	assert.Nil(t, err)
	assert.Equal(t, 10, len(streams[0].Messages))

	for _, stream := range streams {
		for _, message := range stream.Messages {
			t.Logf("message: %v\n", message)
		}
	}
}