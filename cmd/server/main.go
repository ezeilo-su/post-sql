package main

import (
	"log"
	"net/http"

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

	server := &http.Server{
		Addr:         config.ServerPort,
		Handler:      router,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}
	log.Fatal(server.ListenAndServe())
}
