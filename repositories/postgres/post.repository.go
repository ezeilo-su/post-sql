package postrepository

import (
	"context"
	models "postsql/database/models"
	"time"

	"github.com/jmoiron/sqlx"
)

const createPostSql = `
INSERT INTO posts (title, content, image, "user", created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
`

type CreatePostParams struct {
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	User      string    `json:"user" db:"user"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Image     string    `json:"image,omitempty" db:"image,omitempty"`
}

type PostRepository interface {
	Create(context.Context, *CreatePostParams) (*models.PostModel, error)
}

// PostRepositoryImpl handles interactions with the posts table
type PostRepositoryImpl struct {
	db *sqlx.DB
}

// NewPostRepository creates a new type of PostRepository
func NewPostRepository(db *sqlx.DB) PostRepository {
	return &PostRepositoryImpl{db: db}
}

// Create creates a new post in the database
func (r *PostRepositoryImpl) Create(ctx context.Context, post *CreatePostParams) (*models.PostModel, error) {
	var d models.PostModel
	err := r.db.QueryRowContext(ctx, createPostSql, post.Title, post.Content, post.Image, post.User, post.CreatedAt, post.UpdatedAt).Scan(
		&d.UID, &d.Title, &d.Content, &d.Image, &d.User, &d.CreatedAt, &d.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &d, nil
}
