package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	c "github.com/sundayezeilo/post-sql/src/config"
	repository "github.com/sundayezeilo/post-sql/src/repositories"
	api "github.com/sundayezeilo/post-sql/src/routes"
)

func main() {
	config := c.Envs
	ctx := context.Background()
	dbpool, err := pgxpool.New(ctx, config.PostgresURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	defer dbpool.Close()

	log.Println("Connected to Postgres")
	pr := repository.NewPostRepository(dbpool)
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
