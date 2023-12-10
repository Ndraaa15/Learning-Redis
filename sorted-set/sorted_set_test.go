package sortedset

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

//Z for sorted set

func TestSortedSet(t *testing.T) {
	client.ZAdd(ctx, "sortedset", redis.Z{Score: 1, Member: "a"})
	client.ZAdd(ctx, "sortedset", redis.Z{Score: 2, Member: "b"})
	client.ZAdd(ctx, "sortedset", redis.Z{Score: 3, Member: "c"})

	assert.Equal(t, int64(3), client.ZCard(ctx, "sortedset").Val())
	assert.Equal(t, []string{"a", "b", "c"}, client.ZRange(ctx, "sortedset", 0, 2).Val())
	assert.Equal(t, "c", client.ZPopMax(ctx, "sortedset").Val()[0].Member)
	assert.Equal(t, "a", client.ZPopMin(ctx, "sortedset").Val()[0].Member)	
}