package main

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	c "github.com/sundayezeilo/post-sql/src/config"
	api "github.com/sundayezeilo/post-sql/src/routes"
)

func main() {
	config := c.GetConfig()
	db, err := sqlx.Connect("postgres", config.PostgresURL)

	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	log.Println("Connected to Postgres")
	dep := api.Dependencies{DB: db}
	router := api.InitRoutes(dep)

	server := &http.Server{
		Addr:         ":" + config.ServerPort,
		Handler:      router,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}
	log.Println("Server listening on port " + config.ServerPort)
	log.Fatal(server.ListenAndServe())
}
