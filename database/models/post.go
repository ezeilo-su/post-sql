package models

import "time"

// PostModel is the data model for a blog post
type PostModel struct {
	UID       string    `db:"id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	User      string    `db:"user"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Image     string    `db:"image,omitempty"`
}
