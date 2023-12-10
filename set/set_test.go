package set

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

//S for set

func TestSet(t *testing.T) {
	client.SAdd(ctx, "set", "a")
	client.SAdd(ctx, "set", "b")
	client.SAdd(ctx, "set", "c")

	assert.Equal(t, int64(3), client.SCard(ctx, "set").Val())
	assert.Equal(t, true, client.SIsMember(ctx, "set", "a").Val())
	assert.Equal(t, true, client.SIsMember(ctx, "set", "b").Val())
	assert.Equal(t, true, client.SIsMember(ctx, "set", "c").Val())
	assert.Equal(t, false, client.SIsMember(ctx, "set", "d").Val())
	assert.Equal(t, []string{"a", "b", "c"}, client.SMembers(ctx, "set").Val())

	client.SRem(ctx, "set", "a")
	assert.Equal(t, int64(2), client.SCard(ctx, "set").Val())

	client.Del(ctx, "set")
}