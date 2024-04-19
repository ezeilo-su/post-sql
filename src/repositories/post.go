package repositories

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sundayezeilo/post-sql/src/models"
)

const createPostSql = `
INSERT INTO posts (title, content, image, "user", created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
`

type CreatePostParams struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	User      string    `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Image     string    `json:"image,omitempty"`
}

type PostRepository interface {
	Create(context.Context, *CreatePostParams) (*models.PostModel, error)
}

// PostRepositoryImpl handles interactions with the posts table
type PostRepositoryImpl struct {
	dbClient *pgxpool.Pool
}

// NewPostRepository creates a new type of PostRepository
func NewPostRepository(db *pgxpool.Pool) PostRepository {
	return &PostRepositoryImpl{dbClient: db}
}

// Create creates a new post in the database
func (r *PostRepositoryImpl) Create(ctx context.Context, post *CreatePostParams) (*models.PostModel, error) {
	var d models.PostModel
	err := r.dbClient.QueryRow(ctx, createPostSql, post.Title, post.Content, post.Image, post.User, post.CreatedAt, post.UpdatedAt).Scan(
		&d.UID, &d.Title, &d.Content, &d.Image, &d.User, &d.CreatedAt, &d.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &d, nil
}
