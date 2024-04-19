package services

import (
	"context"
	"time"

	"github.com/sundayezeilo/post-sql/src/repositories"
)

type CreatePostParams struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	User      string    `json:"user"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Image     string    `json:"image,omitempty"`
}

type Post struct {
	UID       string    `json:"uid"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	User      string    `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Image     string    `json:"image,omitempty"`
}

type PostService interface {
	CreatePost(context.Context, *CreatePostParams) (Post, error)
}

// PostServiceImpl is the implementation of the PostService interface
type PostServiceImpl struct {
	postRepo repositories.PostRepository
}

// NewPostService creates a new PostService type
func NewPostService(postRepo repositories.PostRepository) PostService {
	return &PostServiceImpl{postRepo: postRepo}
}

// CreatePost handles business logic for creating a new post
func (ps *PostServiceImpl) CreatePost(ctx context.Context, p *CreatePostParams) (Post, error) {
	postDto := &repositories.CreatePostParams{
		Title:     p.Title,
		Content:   p.Content,
		User:      p.User,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Image:     p.Image,
	}

	newPost, err := ps.postRepo.Create(ctx, postDto)

	if err != nil {
		return Post{}, err
	}

	return Post{
		UID:       newPost.UID,
		Title:     newPost.Title,
		Content:   newPost.Content,
		User:      newPost.User,
		CreatedAt: newPost.CreatedAt,
		UpdatedAt: newPost.UpdatedAt,
		Image:     newPost.Image,
	}, nil
}
