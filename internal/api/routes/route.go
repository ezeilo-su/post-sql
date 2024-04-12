package api

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	api "github.com/sundayezeilo/post-spql/internal/api/handlers"
	"github.com/sundayezeilo/post-spql/internal/api/middleware"
	repositories "github.com/sundayezeilo/post-spql/internal/repositories/postgres"
	"github.com/sundayezeilo/post-spql/internal/services"
)

type Dependencies struct {
	DB *sqlx.DB
}

func CreateRoutes(dep Dependencies) *http.ServeMux {
	ps := services.NewPostService(repositories.NewPostRepository(dep.DB))
	pc := api.NewPostHandler(ps)

	router := http.NewServeMux()
	router.HandleFunc("POST /posts", middleware.Logger(pc.CreatePost))

	return router
}
