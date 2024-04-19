package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	c "github.com/sundayezeilo/post-sql/src/config"
	api "github.com/sundayezeilo/post-sql/src/routes"
)

func main() {
	config := c.GetConfig()
	ctx := context.Background()
	dbpool, err := pgxpool.New(ctx, config.PostgresURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	defer dbpool.Close()

	log.Println("Connected to Postgres")
	dep := api.Dependencies{DB: dbpool, Mux: http.NewServeMux()}
	router := api.AddRoutes(dep)

	server := &http.Server{
		Addr:         ":" + config.ServerPort,
		Handler:      router,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}
	log.Println("Server listening on port " + config.ServerPort)
	log.Fatal(server.ListenAndServe())
}
