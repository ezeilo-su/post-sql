package dto

type CreatePostRequest struct {
	User    string `json:"user" validate:"required,min=3,max=100"`
	Title   string `json:"title" validate:"required,min=3,max=100"`
	Content string `json:"content" validate:"required,min=5"`
	Image   string `json:"image,omitempty"`
}

type PostResponse struct {
	ID        string `json:"id"`
	User      string `json:"user"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Image     string `json:"image,omitempty"`
}
