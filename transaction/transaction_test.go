package transaction

import (
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

// we can use transaction with command multi and commit
// but redis using maintain connection pool in internal, so its hard to use MULTI and COMMIT
// we need to use pipeline instead (fn TxPipeline)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   1,
})

var ctx = context.Background()

func TestTransaction(t *testing.T) {
	_, err := client.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.SetEx(ctx, "username", "cupcakezzz", time.Second*10)
		pipe.SetEx(ctx, "email", "indrabrata599@gmail.com", time.Second*10)
		return nil
	})

	assert.Nil(t, err)
	assert.Equal(t, "cupcakezzz", client.Get(ctx, "username").Val())
	assert.Equal(t, "indrabrata599@gmail.com", client.Get(ctx, "email").Val())
}