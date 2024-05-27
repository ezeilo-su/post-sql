package api

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/sundayezeilo/post-sql/api/dto"
	"github.com/sundayezeilo/post-sql/internal/conversion"
	"github.com/sundayezeilo/post-sql/internal/services"
)

// PostHandler is the controller for post resource
type PostHandler struct {
	service   services.PostService
	validator *validator.Validate
}

// NewPostHandler creates a new PostHandler type
func NewPostHandler(service services.PostService) *PostHandler {
	return &PostHandler{
		service:   service,
		validator: validator.New(),
	}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var createPostReq dto.CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&createPostReq); err != nil {
		slog.Error("Error decoding create post payload", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(createPostReq); err != nil {
		slog.Error("Error validating create post payload", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post := conversion.ConvertCreatePostRequestToPost(&createPostReq)

	if err := h.service.CreatePost(context.Background(), post); err != nil {
		slog.Error("Error creating new post", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := conversion.ConvertPostToCreatePostResponse(post)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
