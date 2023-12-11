package geopoint

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

// GeoPoint is a struct that contains the name of the location and the latitude and longitude of the location
func TestGeoPoint(t *testing.T){
	client.GeoAdd(ctx, "location", &redis.GeoLocation{
		Name: "Jakarta",
		Longitude: 106.865036,
		Latitude: -6.175110,
	})

	client.GeoAdd(ctx, "location", &redis.GeoLocation{
		Name: "Bandung",
		Longitude: 107.61861,
		Latitude: -6.90389,
	})


	location := client.GeoPos(ctx, "location", "Jakarta", "Bandung").Val()
	t.Log(location)

	// find the location within a radius of 100 km from Jakarta
	locations := client.GeoSearch(ctx, "location", &redis.GeoSearchQuery{
			Longitude: 106.865036,
			Latitude: -6.175110,
			Radius: 100000,
			RadiusUnit: "km",
	}).Val()

	assert.Equal(t, []string{"Jakarta", "Bandung"}, locations)

	// find the distance between Jakarta and Bandung
	t.Log(client.GeoDist(ctx, "location", "Jakarta", "Bandung", "km").Val())

}