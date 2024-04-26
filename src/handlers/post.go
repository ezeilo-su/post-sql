package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/sundayezeilo/post-sql/src/services"
)

// PostHandler is the controller for post resource
type PostHandler struct {
	service service.PostService
}

// NewPostHandler creates a new PostHandler type
func NewPostHandler(service service.PostService) *PostHandler {
	return &PostHandler{service}
}

// CreatePost handles the incoming POST request to create a new post
func (pc *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var p *service.Post
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	p, err = pc.service.CreatePost(ctx, p)
	if err != nil {
		slog.Error("Error creating new post", err)
		http.Error(w, "Could not create post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}
