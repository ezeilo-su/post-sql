-- name: CreatePost :one
INSERT INTO posts (
  title, content, image, "user", created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;
