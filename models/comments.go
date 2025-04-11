package models

type Comment struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	Username  string `json:"username"`
	Timestamp string `json:"timestamp"`
}