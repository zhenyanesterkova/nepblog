package resolver

import (
	"context"

	"github.com/google/uuid"

	"github.com/zhenyanesterkova/nepblog/internal/app/logger"
	"github.com/zhenyanesterkova/nepblog/internal/feature/comment"
	"github.com/zhenyanesterkova/nepblog/internal/feature/post"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type Repositorie interface {
	Ping() error
	Close() error
	FetchPosts(ctx context.Context, ids []uuid.UUID) ([]post.Post, error)
	FetchCommentsByPostID(ctx context.Context, postID []uuid.UUID) ([]comment.Comment, error)
}

type Resolver struct {
	Repo   Repositorie
	Logger logger.LogrusLogger
}

func NewResolver(
	rep Repositorie,
	log logger.LogrusLogger,
) *Resolver {
	return &Resolver{
		Repo:   rep,
		Logger: log,
	}
}
