package list

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   0,
})

var ctx = context.Background()

func TestString(t *testing.T) {
	client.RPush(ctx, "numbers", 1)
	client.RPush(ctx, "numbers", 2)
	client.RPush(ctx, "numbers", 3)

	assert.Equal(t, int64(3), client.LPop(ctx, "numbers").Val())
	assert.Equal(t, int64(2), client.LPop(ctx, "numbers").Val())
	assert.Equal(t, int64(1), client.LPop(ctx, "numbers").Val())
	

	client.Del(ctx, "numbers")
}
