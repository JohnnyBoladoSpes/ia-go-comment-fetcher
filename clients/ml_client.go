package clients

import (
	"fmt"

	"ia-go-comment-fetcher/models"

	"github.com/go-resty/resty/v2"
)

const SourceInstagram = "Instagram"

type MLClient struct{}

func NewMLClient() *MLClient {
	return &MLClient{}
}

func (c *MLClient) PushCommentsForAnalysis(comments []models.Comment) error {
	client := resty.New()

	var payload []map[string]interface{}
	for _, comment := range comments {
		payload = append(payload, map[string]interface{}{
			"comment_id":   comment.CommentID,
			"user_id":      comment.UserID,
			"text":         comment.Text,
			"media_id":     comment.MediaID,
			"business_id":  comment.BusinessID,
			"source":       SourceInstagram,
		})
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post("http://localhost:9000/api/v1/analysis/analyze")

	if err != nil {
		return err
	}

	if resp.StatusCode() >= 400 {
		return fmt.Errorf("FastAPI returned error %d: %s", resp.StatusCode(), resp.String())
	}

	return nil
}