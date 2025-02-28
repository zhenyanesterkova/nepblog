package handlers

import (
	"context"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/google/uuid"

	"github.com/zhenyanesterkova/nepblog/internal/app/logger"
	"github.com/zhenyanesterkova/nepblog/internal/feature/comment"
	"github.com/zhenyanesterkova/nepblog/internal/feature/post"
	"github.com/zhenyanesterkova/nepblog/internal/http/middleware"
)

const (
	TextServerError                   = "Something went wrong... Server error"
	postLoaderByIDMaxBatch        int = 100
	commentLoaderByPostIDMaxBatch int = 10
)

type Repositorie interface {
	Ping() error
	Close() error
	FetchPosts(ctx context.Context, ids []uuid.UUID) ([]post.Post, error)
	FetchCommentsByPostID(ctx context.Context, postID []uuid.UUID) ([]comment.Comment, error)
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
	postLoader := post.NewConfiguredLoaderByID(rh.Repo, postLoaderByIDMaxBatch)
	commentLoader := comment.NewConfiguredLoaderByPostID(rh.Repo, commentLoaderByPostIDMaxBatch)

	mdlWare := middleware.NewMiddlewareStruct(rh.Logger, postLoader, commentLoader)
	router.Use(mdlWare.RequestLogger)
	router.Use(mdlWare.GZipMiddleware)
	router.Route("/", func(r chi.Router) {
		r.Handle("/api", rh.GraphQLHandler())
		r.Get("/", playground.Handler("Posts", "/api"))
	})
}
