package service

import (
	"context"
	"time"

	"github.com/sundayezeilo/post-sql/src/models"
	"github.com/sundayezeilo/post-sql/src/repositories"
)

type PostDto struct {
	UID       string    `json:"uid"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	User      string    `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Image     string    `json:"image,omitempty"`
}
type CreatePostParams struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	User      string    `json:"user"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Image     string    `json:"image,omitempty"`
}

type PostService interface {
	CreatePost(context.Context, *CreatePostParams) (PostDto, error)
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
func (ps *PostServiceImpl) CreatePost(ctx context.Context, p *CreatePostParams) (PostDto, error) {
	postDto := &model.PostModel{
		Title:     p.Title,
		Content:   p.Content,
		User:      p.User,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Image:     p.Image,
	}

	newPost, err := ps.postRepo.Create(ctx, postDto)

	if err != nil {
		return PostDto{}, err
	}

	return PostDto{
		UID:       newPost.UID,
		Title:     newPost.Title,
		Content:   newPost.Content,
		User:      newPost.User,
		CreatedAt: newPost.CreatedAt,
		UpdatedAt: newPost.UpdatedAt,
		Image:     newPost.Image,
	}, nil
}
