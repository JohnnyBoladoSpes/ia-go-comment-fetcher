package db

import (
	"log"
	"os"
	"sync"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client
var once sync.Once

func InitRedis() *redis.Client {
	once.Do(func() {
		addr := os.Getenv("REDIS_HOST")
		if addr == "" {
			addr = "localhost:6379"
		}
		redisClient = redis.NewClient(&redis.Options{
			Addr: addr,
			DB:   0,
		})
	})

	if redisClient == nil {
		log.Fatal("Redis client not initialized")
	}

	return redisClient
}