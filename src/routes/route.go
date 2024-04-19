package api

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	api "github.com/sundayezeilo/post-sql/src/handlers"
	"github.com/sundayezeilo/post-sql/src/middleware"
	"github.com/sundayezeilo/post-sql/src/repositories"
	"github.com/sundayezeilo/post-sql/src/services"
)

type Dependencies struct {
	DB  *pgxpool.Pool
	Mux *http.ServeMux
}

func AddRoutes(dep Dependencies) *http.ServeMux {
	ps := services.NewPostService(repositories.NewPostRepository(dep.DB))
	pc := api.NewPostHandler(ps)

	dep.Mux.HandleFunc("POST /posts", middleware.Logger(pc.CreatePost))

	return dep.Mux
}
