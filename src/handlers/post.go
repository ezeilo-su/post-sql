package api

import (
	"encoding/json"
	"log"
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
	var postParams service.CreatePostParams
	err := json.NewDecoder(r.Body).Decode(&postParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var post service.PostDto
	ctx := r.Context()
	post, err = pc.service.CreatePost(ctx, &postParams)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, "Could not create post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}
