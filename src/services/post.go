package service

import (
	"context"

	"github.com/sundayezeilo/post-sql/src/models"
	"github.com/sundayezeilo/post-sql/src/repositories"
)

type PostService interface {
	CreatePost(context.Context, *model.Post) (*model.Post, error)
}

// PostServiceImpl is the implementation of the PostService interface
type PostServiceImpl struct {
	postRepo repository.PostRepository
}

// NewPostService creates a new PostService type
func NewPostService(postRepo repository.PostRepository) PostService {
	return &PostServiceImpl{postRepo: postRepo}
}

// CreatePost handles business logic for creating a new post
func (ps *PostServiceImpl) CreatePost(ctx context.Context, p *model.Post) (*model.Post, error) {
	pm := &model.Post{
		Title:   p.Title,
		Content: p.Content,
		User:    p.User,
		Image:   p.Image,
	}

	pm, err := ps.postRepo.Create(ctx, pm)

	if err != nil {
		return nil, err
	}

	return pm, nil
}
