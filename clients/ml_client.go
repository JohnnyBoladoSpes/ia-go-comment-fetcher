package clients

import (
	"fmt"

	"ia-go-comment-fetcher/models"

	"github.com/go-resty/resty/v2"
)

type MLClient struct{}

func NewMLClient() *MLClient {
	return &MLClient{}
}

func (c *MLClient) PushCommentsForAnalysis(comments []models.Comment) error {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(comments).
		Post("http://localhost:8000/analyze-comments")

	if err != nil {
		return err
	}

	if resp.StatusCode() >= 400 {
		return fmt.Errorf("FastAPI returned error %d: %s", resp.StatusCode(), resp.String())
	}

	return nil
}