package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/sundayezeilo/post-sql/src/services"
)

// PostHandler is the controller for post resource
type PostHandler struct {
	service services.PostService
}

// NewPostHandler creates a new PostHandler type
func NewPostHandler(service services.PostService) *PostHandler {
	return &PostHandler{service}
}

// CreatePost handles the incoming POST request to create a new post
func (pc *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post services.CreatePostParams
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var newPost *services.Post
	newPost, err = pc.service.CreatePost(context.Background(), &post)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, "Could not create post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(*newPost)
}
