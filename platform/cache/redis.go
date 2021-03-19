package cache

import (
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func RedisConnection() *redis.Client {
	// Define DB number.
	databaseNumber, _ := strconv.Atoi(os.Getenv("REDIS_DB_NUMBER"))

	// Set Redis options.
	options := &redis.Options{
		Addr:     os.Getenv("REDIS_SERVER_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       databaseNumber,
	}

	return redis.NewClient(options)
}
