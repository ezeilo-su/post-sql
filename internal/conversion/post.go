package conversion

import (
	"github.com/sundayezeilo/post-sql/api/dto"
	"github.com/sundayezeilo/post-sql/internal/models"
)

// ConvertCreatePostRequestToPost converts a CreatePostRequest DTO to a Post model.
func ConvertCreatePostRequestToPost(req *dto.CreatePostRequest) *models.Post {
	return &models.Post{
		User:    req.User,
		Title:   req.Title,
		Content: req.Content,
		Image:   req.Image,
	}
}

// ConvertPostToCreatePostResponse converts a Post model to a CreatePostResponse DTO.
func ConvertPostToCreatePostResponse(post *models.Post) *dto.CreatePostResponse {
	return &dto.CreatePostResponse{
		ID:        post.ID,
		User:      post.User,
		Title:     post.Title,
		Content:   post.Content,
		Image:     post.Image,
		CreatedAt: post.CreatedAt.String(),
		UpdatedAt: post.UpdatedAt.String(),
	}
}
