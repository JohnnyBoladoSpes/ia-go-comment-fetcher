package routes

import (
	"net/http"

	"ia-go-comment-fetcher/controllers"

	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(mongoClient *mongo.Client) {
	controller := controllers.NewCommentController(mongoClient)
	http.HandleFunc("/fetch-comments", controller.FetchComments)
}