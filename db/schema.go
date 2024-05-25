package database

import "time"

// PostDB represents the database schema for the posts table
type PostDB struct {
	UID       string    `db:"id"`
	User      string    `db:"user"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Image     string    `db:"image,omitempty"`
}
