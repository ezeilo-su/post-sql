package api

import (
	"context"
	"log"
	"net/http"
	"time"

	api "github.com/sundayezeilo/post-sql/src/handlers"
	"github.com/sundayezeilo/post-sql/src/middleware"
	repository "github.com/sundayezeilo/post-sql/src/repositories"
	service "github.com/sundayezeilo/post-sql/src/services"
)

type APIServer struct {
	Addr         string
	Ctx          context.Context
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Repository   *repository.Repository
}

// func NewAPIServer(cfg *APIServer) *APIServer {
// 	return &APIServer{
// 		ctx:          cfg.ctx,
// 		addr:         cfg.addr,
// 		repository:   cfg.repository,
// 		ReadTimeout:  cfg.ReadTimeout,
// 		WriteTimeout: cfg.WriteTimeout,
// 	}
// }

func (s *APIServer) Run() {
	ps := service.NewPostService(s.Ctx, s.Repository.Post)
	pc := api.NewPostHandler(ps)

	mux := http.NewServeMux()
	apiV1Mux := http.NewServeMux()
	apiV1Mux.HandleFunc("POST /posts", middleware.Logger(pc.CreatePost))

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiV1Mux))

	httpServer := &http.Server{
		Addr:         ":" + s.Addr,
		Handler:      mux,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
	}

	go func() {
		log.Printf("listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error listening and serving: %s\n", err)
		}
	}()

	// Block the main goroutine (optional, based on shutdown logic)
	select {}
}
