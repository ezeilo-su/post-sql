package models

import "time"

// Post is the data model for a blog post
type PostModel struct {
	UID       int       `db:"id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	User      string    `db:"user"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Image     string    `db:"image,omitempty"`
}
