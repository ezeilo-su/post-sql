package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/sundayezeilo/post-spql/internal/api/handlers"
	c "github.com/sundayezeilo/post-spql/internal/config"
	"github.com/sundayezeilo/post-spql/internal/repositories/postgres"
	"github.com/sundayezeilo/post-spql/internal/services"
)

func main() {
	config := c.GetConfig()
	db, err := sqlx.Connect("postgres", config.PostgresURL)

	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()
	log.Println("Connected to Postgres")

	ps := services.NewPostService(repositories.NewPostRepository(db))
	pc := api.NewPostHandler(ps)

	router := http.NewServeMux()
	router.HandleFunc("POST /posts", pc.CreatePost)

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
