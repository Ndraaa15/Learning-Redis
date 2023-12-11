package pubsub

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   1,
})

var ctx = context.Background()

func TestSubscribePubSub (t *testing.T){
	pubSub := client.Subscribe(ctx, "channel-1")
	// pubSub is a *redis.PubSub

	for i := 0; i < 10; i++ {
		message, _ := pubSub.ReceiveMessage(ctx)
		// message is a *redis.Message
		t.Log(message.Payload)
	}
	pubSub.Close()
}

func TestPublishPubSub (t *testing.T){
	for i:= 0; i < 10; i++ {
		err := client.Publish(ctx, "channel-1", "hello world").Err()
		assert.Nil(t, err)
	}
}