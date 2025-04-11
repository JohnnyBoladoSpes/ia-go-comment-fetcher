package clients

import (
	"encoding/json"
	"os"

	"ia-go-comment-fetcher/models"
)

type InstagramClient struct{}

func NewInstagramClient() *InstagramClient {
	return &InstagramClient{}
}

func (c *InstagramClient) GetCommentsMock(mediaID string) ([]models.Comment, error) {
	data, err := os.ReadFile("assets/mock_comments.json")
	if err != nil {
		return nil, err
	}

	var comments []models.Comment
	if err := json.Unmarshal(data, &comments); err != nil {
		return nil, err
	}

	return comments, nil
}