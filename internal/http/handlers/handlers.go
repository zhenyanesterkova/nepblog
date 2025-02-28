package handlers

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"

	"github.com/zhenyanesterkova/nepblog/internal/app/logger"
	"github.com/zhenyanesterkova/nepblog/internal/http/middleware"
)

const (
	TextServerError = "Something went wrong... Server error"
)

type Repositorie interface {
	Ping() error
	Close() error
}

type RepositorieHandler struct {
	Repo   Repositorie
	Logger logger.LogrusLogger
}

func NewRepositorieHandler(
	rep Repositorie,
	log logger.LogrusLogger,
) *RepositorieHandler {
	return &RepositorieHandler{
		Repo:   rep,
		Logger: log,
	}
}

func (rh *RepositorieHandler) InitChiRouter(router *chi.Mux) {
	mdlWare := middleware.NewMiddlewareStruct(rh.Logger)
	router.Use(mdlWare.RequestLogger)
	router.Use(mdlWare.GZipMiddleware)
	router.Route("/", func(r chi.Router) {
		r.Handle("/api", rh.GraphQLHandler())
		r.Get("/", playground.Handler("Posts", "/api"))
	})
}
