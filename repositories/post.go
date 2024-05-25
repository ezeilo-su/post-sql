package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sundayezeilo/post-sql/db"
)

const createPostSql = `
INSERT INTO posts (title, content, image, "user", created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
`

type PostDto struct {
	UID       string    `json:"uid"`
	User      string    `json:"user"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Image     string    `json:"image,omitempty"`
}

type PostRepository interface {
	Create(context.Context, *PostDto) (*PostDto, error)
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
func (r *PostRepositoryImpl) Create(ctx context.Context, p *PostDto) (*PostDto, error) {
	postDb := &database.PostDB{
		User:      p.User,
		Title:     p.Title,
		Content:   p.Content,
		Image:     p.Image,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := r.dbClient.QueryRow(ctx, createPostSql, postDb.Title, postDb.Content, postDb.Image, postDb.User, postDb.CreatedAt, postDb.UpdatedAt).Scan(&p.UID, &p.Title, &p.Content, &p.Image, &p.User, &p.CreatedAt, &p.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return p, nil
}
