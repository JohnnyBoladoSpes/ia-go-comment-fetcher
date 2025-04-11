package usecases

import (
	"errors"

	"ia-go-comment-fetcher/clients"
	"ia-go-comment-fetcher/models"
	"ia-go-comment-fetcher/services"
)

var cacheService = services.NewRequestCacheService()
var dataService = services.NewRequestDataService()
var igClient = clients.NewInstagramClient()
var mlClient = clients.NewMLClient()

func FetchCommentsUseCase(mediaID string, businessID string) ([]models.Comment, error) {
	if ts, err := cacheService.GetLastRequestTimestamp(mediaID, businessID); err == nil && ts != "" {
		return nil, errors.New("request already made recently")
	}

	if err := cacheService.SaveRequestTimestamp(mediaID, businessID); err != nil {
		return nil, err
	}

	comments, err := igClient.GetCommentsMock(mediaID)
	if err != nil {
		return nil, err
	}

	if err := mlClient.PushCommentsForAnalysis(comments); err != nil {
		return nil, err
	}

	if err := dataService.StoreRequest(mediaID, businessID); err != nil {
		return nil, err
	}

	return comments, nil
}