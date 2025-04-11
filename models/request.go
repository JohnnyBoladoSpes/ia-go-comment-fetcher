package models

type FetchCommentsRequest struct {
	MediaId string `json:"media_id" validate:"required"`
	BusinessId string `json:"business_id" validate:"required"`
}