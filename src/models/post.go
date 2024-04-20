package model

import "time"

// PostModel is the data model for a blog post
type Post struct {
	UID       string    `json:"uid"`
	User      string    `json:"user"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Image     string    `json:"image,omitempty"`
}
