package services

import (
	"context"
	"time"

	"ia-go-comment-fetcher/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type RequestDataService struct {
	mongoClient *mongo.Client
}

func NewRequestDataService(client *mongo.Client) *RequestDataService {
	return &RequestDataService{
		mongoClient: client,
	}
}

func (service *RequestDataService) StoreRequest(mediaID, businessID string) error {
	collection := service.mongoClient.Database("comment_fetcher").Collection("requests")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := models.RequestLog{
		MediaId:     mediaID,
		BusinessId:  businessID,
		RequestedAt: time.Now(),
	}

	_, err := collection.InsertOne(ctx, request)
	return err
}