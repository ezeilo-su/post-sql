package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// Post is the data model for a blog post
type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Image     string    `json:"image,omitempty"`
	User      string    `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// PostService is the interface for the CRUD operations on posts
type PostService interface {
	CreatePost(post Post) (Post, error)
}

// PostServiceImpl is the implementation of the PostService interface
type PostServiceImpl struct {
	db *sql.DB
}

// NewPostService creates a new PostService object
func NewPostService(db *sql.DB) PostService {
	return &PostServiceImpl{db}
}

func createTable(client *sql.DB, name string) error {
	_, err := client.QueryContext(
		context.Background(),
		"CREATE TABLE IF NOT EXISTS "+name+` (
    id SERIAL PRIMARY KEY,
    title VARCHAR(300),
    content VARCHAR(1000),
    image VARCHAR(255),
    "user" VARCHAR(50),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
	);`)
	return err
}

// CreatePost inserts a new post into the database
func (ps *PostServiceImpl) CreatePost(post Post) (Post, error) {

	err := createTable(ps.db, "posts")

	if err != nil {
		return Post{}, err
	}

	query := `
		INSERT INTO posts (title, content, image, "user", created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`
	var id int
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
	err = ps.db.QueryRowContext(
		context.Background(),
		query,
		post.Title, post.Content, post.Image, post.User, post.CreatedAt, post.UpdatedAt,
	).Scan(&id)
	if err != nil {
		return Post{}, err
	}
	post.ID = id
	return post, nil
}

// PostController is the controller for the CRUD endpoints
type PostController struct {
	service PostService
}

// NewPostController creates a new PostController object
func NewPostController(service PostService) *PostController {
	return &PostController{service}
}

// CreatePost handles the POST request to create a new post
func (pc *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var newPost Post
	newPost, err = pc.service.CreatePost(post)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, "Could not create post", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPost)
}

func main() {
	dbURI := os.Getenv("PG_URI")
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Connected to Postgres")

	// Create a service and a controller for posts
	service := NewPostService(db)
	controller := NewPostController(service)

	// Create a router and define the routes for posts
	router := http.NewServeMux()
	router.HandleFunc("POST /posts", controller.CreatePost)
	router.Handle("/", http.NotFoundHandler())

	// Run the server
	server := &http.Server{
		Addr:         ":80",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
