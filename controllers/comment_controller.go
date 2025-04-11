package controllers

import (
	"encoding/json"
	"net/http"

	"ia-go-comment-fetcher/models"
	"ia-go-comment-fetcher/usecases"
	"ia-go-comment-fetcher/utils"
)

type CommentController struct{}

func NewCommentController() *CommentController {
	return &CommentController{}
}

func (cc *CommentController) FetchComments(w http.ResponseWriter, r *http.Request) {
	var req models.FetchCommentsRequest

	if err := utils.ValidateRequest(w, r, &req); err != nil {
		return
	}

	comments, err := usecases.FetchCommentsUseCase(req.MediaId, req.BusinessId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comments)
}