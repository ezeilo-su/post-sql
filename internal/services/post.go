package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sundayezeilo/post-sql/internal/models"
	repository "github.com/sundayezeilo/post-sql/internal/repositories"
)

type PostService interface {
	CreatePost(ctx context.Context, post *models.Post) error
}

// postService is the implementation of the PostService interface
type postService struct {
	ctx      context.Context
	postRepo repository.PostRepository
}

// NewPostService creates a new PostService type
func NewPostService(ctx context.Context, postRepo repository.PostRepository) PostService {
	return &postService{ctx, postRepo}
}

// CreatePost handles business logic for creating a new post
func (ps *postService) CreatePost(ctx context.Context, pm *models.Post) error {
	pm.ID = uuid.New().String()
	pm.CreatedAt = time.Now()
	pm.UpdatedAt = pm.CreatedAt

	err := ps.postRepo.Create(ctx, pm)

	if err != nil {
		return err
	}

	return nil
}
