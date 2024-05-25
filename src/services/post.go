package service

import (
	"context"
	"time"

	repository "github.com/sundayezeilo/post-sql/src/repositories"
)

type Post struct {
	UID       string    `json:"uid"`
	User      string    `json:"user"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Image     string    `json:"image,omitempty"`
}

type PostService interface {
	CreatePost(context.Context, *Post) (*Post, error)
}

// PostServiceImpl is the implementation of the PostService interface
type PostServiceImpl struct {
	ctx      context.Context
	postRepo repository.PostRepository
}

// NewPostService creates a new PostService type
func NewPostService(ctx context.Context, postRepo repository.PostRepository) PostService {
	return &PostServiceImpl{ctx, postRepo}
}

// CreatePost handles business logic for creating a new post
func (ps *PostServiceImpl) CreatePost(ctx context.Context, p *Post) (*Post, error) {
	pm := &repository.PostDto{
		Title:   p.Title,
		Content: p.Content,
		User:    p.User,
		Image:   p.Image,
	}
	var newPost *repository.PostDto
	newPost, err := ps.postRepo.Create(ctx, pm)

	if err != nil {
		return nil, err
	}

	p.UID = newPost.UID
	p.User = newPost.User
	p.Title = newPost.Title
	p.Content = newPost.Content
	p.CreatedAt = newPost.CreatedAt
	p.UpdatedAt = newPost.UpdatedAt
	p.Image = newPost.Image

	return p, nil
}
