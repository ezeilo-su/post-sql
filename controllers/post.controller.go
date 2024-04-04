package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	services "postsql/services"
)

// PostController is the controller for post resource
type PostController struct {
	service services.PostService
}

// NewPostController creates a new PostController type
func NewPostController(service services.PostService) *PostController {
	return &PostController{service}
}

// CreatePost handles the incoming POST request to create a new post
func (pc *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
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
