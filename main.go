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

	server := &http.Server{
		Addr:         ":80",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
