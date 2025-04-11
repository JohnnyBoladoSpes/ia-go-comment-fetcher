package routes

import (
	"net/http"

	"ia-go-comment-fetcher/controllers"
)

func RegisterRoutes() {
	controller := controllers.NewCommentController()
	http.HandleFunc("/fetch-comments", controller.FetchComments)
}