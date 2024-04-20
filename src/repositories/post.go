package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sundayezeilo/post-sql/src/models"
)

const createPostSql = `
INSERT INTO posts (title, content, image, "user", created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
`

type PostRepository interface {
	Create(context.Context, *model.PostModel) (*model.PostModel, error)
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
func (r *PostRepositoryImpl) Create(ctx context.Context, p *model.PostModel) (*model.PostModel, error) {
	err := r.dbClient.QueryRow(ctx, createPostSql, p.Title, p.Content, p.Image, p.User, p.CreatedAt, p.UpdatedAt).Scan(&p.UID)

	if err != nil {
		return nil, err
	}

	return p, nil
}
