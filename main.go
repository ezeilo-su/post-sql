package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	controllers "postsql/controllers"
	repositories "postsql/repositories/postgres"
	services "postsql/services"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	// db, err := sqlx.Open("postgres", dbURL)
	db, err := sqlx.Connect("postgres", dbURL)

	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()
	log.Println("Connected to Postgres")

	service := services.NewPostService(repositories.NewPostRepository(db))
	// Create a service and a controller for posts
	controller := controllers.NewPostController(service)

	// Create a router and define the routes for posts
	router := http.NewServeMux()
	router.HandleFunc("POST /posts", controller.CreatePost)
	router.Handle("/", http.NotFoundHandler())

	server := &http.Server{
		Addr:         ":80",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
