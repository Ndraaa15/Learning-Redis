package hyperloglog

import (
	"context"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/redis/go-redis/v9"
)



var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   0,
})

var ctx = context.Background()

// HyperLogLog is to count the unique value

func TestHyperLogLog(t *testing.T){
	client.PFAdd(ctx, "alphabet", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j")
	client.PFAdd(ctx, "alphabet", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j")
	client.PFAdd(ctx, "alphabet", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t")

	assert.Equal(t, int64(19), client.PFCount(ctx, "alphabet").Val())
}