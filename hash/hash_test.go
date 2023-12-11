package hash

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

// H for hash
func TestHash (t *testing.T){
	client.HSet(ctx, "user:1", "username", "cupcakezzz")
	client.HSet(ctx, "user:1", "password", "123456")
	client.HSet(ctx, "user:1", "email", "indrabrata599@gmail.com")

	user := client.HGetAll(ctx, "user:1").Val()
	assert.Equal(t, "cupcakezzz", user["username"])
	assert.Equal(t, "123456", user["password"])
	assert.Equal(t, "indrabrata599@gmail.com", user["email"])
	
	client.HDel(ctx, "user:1")
}