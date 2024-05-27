package main

import (
	"context"
	"fmt"
	"log"
	"os"

	api "github.com/sundayezeilo/post-sql/api/routes"
	c "github.com/sundayezeilo/post-sql/config"
	"github.com/sundayezeilo/post-sql/internal/db"
	repository "github.com/sundayezeilo/post-sql/internal/repositories"
)

func main() {
	config := c.Envs
	ctx := context.Background()

	store, err := db.NewPostgresDB(ctx, config.PostgresURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create DB connection: %v\n", err)
		os.Exit(1)
	}

	defer store.Close()

	log.Println("Connected to Postgres")
	pr := repository.NewPostRepository(store)
	repos := &repository.Repository{Post: pr}
	apiServer := &api.APIServer{
		Addr:         config.ServerPort,
		Ctx:          ctx,
		Repository:   repos,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}

	apiServer.Run()
}
