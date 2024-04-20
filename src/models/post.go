package model

import "time"

// PostModel is the data model for a blog post
type Post struct {
	UID       string    `db:"id" json:"uid"`
	User      string    `db:"user" json:"user"`
	Title     string    `db:"title" json:"title"`
	Content   string    `db:"content" json:"content"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Image     string    `db:"image,omitempty" json:"image,omitempty"`
}
