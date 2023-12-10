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


func TestList(t *testing.T) {
	client.RPush(ctx, "numbers", 1)
	client.RPush(ctx, "numbers", 2)
	client.RPush(ctx, "numbers", 3)

	assert.Equal(t, "1", client.LPop(ctx, "numbers").Val())
	assert.Equal(t, "2", client.LPop(ctx, "numbers").Val())
	assert.Equal(t, "3", client.LPop(ctx, "numbers").Val())
	
	client.Del(ctx, "numbers")
}

//RPush is a function that pushes a value to the right of a list
//LPop is a function that pops a value from the left of a list