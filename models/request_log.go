package models

import "time"

type RequestLog struct {
	MediaId string `bson:"media_id"`
	BusinessId string `bson:"business_id"`
	RequestedAt time.Time `bson:"requested_at"`
}