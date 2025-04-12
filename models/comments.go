package models

type Comment struct {
	CommentID  string `json:"comment_id"`
	UserID     string `json:"user_id"`
	Source     string `json:"source,omitempty"`
	Text       string `json:"text"`
	RequestedAt string `json:"requested_at,omitempty"`
	MediaID    string `json:"media_id"`
	BusinessID string `json:"business_id"`
}