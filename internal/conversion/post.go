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

// ConvertPostToPostResponse converts a Post model to a PostResponse DTO.
func ConvertPostToPostResponse(post *models.Post) *dto.PostResponse {
	return &dto.PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		User:      post.User,
		Content:   post.Content,
		Image:     post.Image,
		CreatedAt: post.CreatedAt.Format("2009-11-10 23:00:00"),
		UpdatedAt: post.UpdatedAt.Format("2009-11-10 23:00:00"),
	}
}
