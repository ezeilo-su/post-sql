package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sundayezeilo/post-sql/api/routes"
	"github.com/sundayezeilo/post-sql/config"
	"github.com/sundayezeilo/post-sql/internal/db"
	"github.com/sundayezeilo/post-sql/internal/repositories"
)

func main() {
	cfg := config.LoadEnv()
	ctx := context.Background()

	store, err := db.NewPostgresDB(ctx, cfg.PostgresURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create DB connection: %v\n", err)
		os.Exit(1)
	}

	defer store.Close()

	log.Println("Connected to Postgres")
	pr := repositories.NewPostRepository(store)
	repos := &repositories.Repository{Post: pr}
	server := &api.Server{
		Addr:         cfg.ServerPort,
		Ctx:          ctx,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	httpServer := server.AddRoutes(repos)

	go func() {
		log.Printf("listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error listening and serving: %s\n", err)
		}
	}()

	// Block the main goroutine (optional, based on shutdown logic)
	select {}
}
