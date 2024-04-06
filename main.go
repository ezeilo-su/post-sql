package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/sundayezeilo/post-spql/controllers"
	"github.com/sundayezeilo/post-spql/repositories/postgres"
	"github.com/sundayezeilo/post-spql/services"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatalln("DATABASE_URL not set in the ENV")
	}
	db, err := sqlx.Connect("postgres", dbURL)

	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()
	log.Println("Connected to Postgres")

	ps := services.NewPostService(repositories.NewPostRepository(db))
	pc := controllers.NewPostController(ps)

	router := http.NewServeMux()
	router.HandleFunc("POST /posts", pc.CreatePost)
	router.Handle("/", http.NotFoundHandler())

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Fatalln("SERVER_PORT must be set in the ENV")
	}

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
