package controllers

import (
	"encoding/json"
	"net/http"

	"ia-go-comment-fetcher/models"
	"ia-go-comment-fetcher/usecases"
	"ia-go-comment-fetcher/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

type CommentController struct {
	dbClient *mongo.Client
}

func NewCommentController(client *mongo.Client) *CommentController {
	return &CommentController{
		dbClient: client,
	}
}

func (cc *CommentController) FetchComments(w http.ResponseWriter, r *http.Request) {
	var req models.FetchCommentsRequest

	if err := utils.ValidateRequest(w, r, &req); err != nil {
		return
	}

	usecase := usecases.FetchCommentsUseCase(cc.dbClient)
	comments, err := usecase.Fetch(req.MediaId, req.BusinessId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comments)
}