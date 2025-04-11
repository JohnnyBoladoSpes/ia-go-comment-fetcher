package usecases

import (
	"errors"
	"ia-go-comment-fetcher/clients"
	"ia-go-comment-fetcher/models"
	"ia-go-comment-fetcher/services"

	"go.mongodb.org/mongo-driver/mongo"
)
type CommentFetcher struct {
	CacheService *services.RequestCacheService
	DataService  *services.RequestDataService
	IgClient     *clients.InstagramClient
	MlClient     *clients.MLClient
}

func FetchCommentsUseCase(
	client *mongo.Client,
) *CommentFetcher {
	return &CommentFetcher{
		CacheService: services.NewRequestCacheService(),
		DataService:  services.NewRequestDataService(client),
		IgClient:     clients.NewInstagramClient(),
		MlClient:     clients.NewMLClient(),
	}
}

func (f *CommentFetcher) Fetch(mediaID, businessID string) ([]models.Comment, error) {
	if ts, err := f.CacheService.GetLastRequestTimestamp(mediaID, businessID); err == nil && ts != "" {
		return nil, errors.New("request already made recently")
	}

	comments, err := f.IgClient.GetCommentsMock(mediaID)
	if err != nil {
		return nil, err
	}

	if err := f.DataService.StoreRequest(mediaID, businessID); err != nil {
		return nil, err
	}

	if err := f.CacheService.SaveRequestTimestamp(mediaID, businessID); err != nil {
		return nil, err
	}

	if err := f.MlClient.PushCommentsForAnalysis(comments); err != nil {
		return nil, err
	}

	return comments, nil
}