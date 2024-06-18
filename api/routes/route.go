package api

import (
	"context"
	"net/http"
	"time"

	"github.com/sundayezeilo/post-sql/api/handlers"
	"github.com/sundayezeilo/post-sql/internal/repositories"
	"github.com/sundayezeilo/post-sql/internal/services"
)

type Server struct {
	Addr         string
	Ctx          context.Context
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (s *Server) AddRoutes(repository *repositories.Repository) *http.Server {
	ps := services.NewPostService(s.Ctx, repository.Post)
	ph := api.NewPostHandler(ps)

	apiV1Mux := http.NewServeMux()
	ph.RegisterRoutes(apiV1Mux)

	mux := http.NewServeMux()
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiV1Mux))

	httpServer := &http.Server{
		Addr:         ":" + s.Addr,
		Handler:      mux,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
	}
	return httpServer
}
