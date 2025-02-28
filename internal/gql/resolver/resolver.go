package resolver

import "github.com/zhenyanesterkova/nepblog/internal/app/logger"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type Repositorie interface {
	Ping() error
	Close() error
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
