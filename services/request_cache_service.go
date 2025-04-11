package services

import (
	"context"
	"time"

	"ia-go-comment-fetcher/db"

	"github.com/redis/go-redis/v9"
)

type RequestCacheService struct {
	redisClient *redis.Client
}

func NewRequestCacheService() *RequestCacheService {
	return &RequestCacheService{
		redisClient: db.InitRedis(),
	}
}

func (service *RequestCacheService) SaveRequestTimestamp(mediaID, businessID string) error {
	key := "comments:" + mediaID + ":" + businessID
	value := time.Now().UTC().Format(time.RFC3339)
	return service.redisClient.Set(context.Background(), key, value, 2*time.Hour).Err()
}

func (service *RequestCacheService) GetLastRequestTimestamp(mediaID, businessID string) (string, error) {
	key := "comments:" + mediaID + ":" + businessID
	return service.redisClient.Get(context.Background(), key).Result()
}