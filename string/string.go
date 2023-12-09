package string

import (
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   0,
})

var ctx = context.Background()

func TestString (t *testing.T) {
	client.SetEx(ctx, "username", "cupcakezzz", 3 * time.Second)

	result, err := client.Get(ctx, "username").Result()
	assert.Nil(t, err)
	assert.Equal(t, "cupcakezzz", result)

	time.Sleep(3 * time.Second)

	_, err = client.Get(ctx, "username").Result()
	assert.NotNil(t, err)
}