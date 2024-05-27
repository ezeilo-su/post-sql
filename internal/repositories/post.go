package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sundayezeilo/post-sql/internal/models"
)

const createPostSql = `
INSERT INTO posts (id, title, content, image, "user", created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;
`

type PostRepository interface {
	Create(ctx context.Context, post *models.Post) error
}

// postRepository handles interactions with the posts table
type postRepository struct {
	db *pgxpool.Pool
}

// NewPostRepository creates a new type of PostRepository
func NewPostRepository(db *pgxpool.Pool) PostRepository {
	return &postRepository{db: db}
}

// Create creates a new post in the database
func (r *postRepository) Create(ctx context.Context, pm *models.Post) error {
	err := r.db.QueryRow(ctx, createPostSql, pm.ID, pm.Title, pm.Content, pm.Image, pm.User, pm.CreatedAt, pm.UpdatedAt).Scan(&pm.ID, &pm.Title, &pm.Content, &pm.Image, &pm.User, &pm.CreatedAt, &pm.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}
