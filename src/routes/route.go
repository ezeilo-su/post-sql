package api

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/sundayezeilo/post-sql/src/handlers"
	"github.com/sundayezeilo/post-sql/src/middleware"
	"github.com/sundayezeilo/post-sql/src/repositories"
	"github.com/sundayezeilo/post-sql/src/services"
)

type Dependencies struct {
	DB *sqlx.DB
}

func InitRoutes(dep Dependencies) *http.ServeMux {
	ps := services.NewPostService(repositories.NewPostRepository(dep.DB))
	pc := api.NewPostHandler(ps)

	router := http.NewServeMux()
	router.HandleFunc("POST /posts", middleware.Logger(pc.CreatePost))

	return router
}
