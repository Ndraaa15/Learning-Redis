package connection

import (
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:  0,
})


func TestConnection(t *testing.T) {
		assert.NotNil(t, client)

		err := client.Close()
		assert.Nil(t, err)
}


