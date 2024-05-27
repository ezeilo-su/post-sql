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
	ctx  context.Context
	repo repository.PostRepository
}

// NewPostService creates a new PostService type
func NewPostService(ctx context.Context, repo repository.PostRepository) PostService {
	return &postService{ctx, repo}
}

// CreatePost handles business logic for creating a new post
func (ps *postService) CreatePost(ctx context.Context, pm *models.Post) error {
	uuidV7, err := uuid.NewV7()
	if err != nil {
		return err
	}

	// Generate a new UUID V7 for the post
	pm.ID = uuidV7.String()
	pm.CreatedAt = time.Now().UTC()
	pm.UpdatedAt = pm.CreatedAt

	err = ps.repo.Create(ctx, pm)

	if err != nil {
		return err
	}

	return nil
}
