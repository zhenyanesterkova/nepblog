package handlers

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/zhenyanesterkova/nepblog/internal/gql/resolver"
	"github.com/zhenyanesterkova/nepblog/internal/gql/runtime"
)

const websocketKeepAlivePingInterval = 5 * time.Second
const queryCacheLRUSize = 1000
const automaticPersistedQueryCacheLRUSize = 100
const complexityLimit = 1000

func (rh *RepositorieHandler) GraphQLHandler() *handler.Server {
	cfg := runtime.Config{Resolvers: resolver.NewResolver(rh.Repo, rh.Logger)}

	handler := handler.New(
		runtime.NewExecutableSchema(
			cfg,
		),
	)

	// Transports
	handler.AddTransport(transport.Websocket{
		KeepAlivePingInterval: websocketKeepAlivePingInterval,
	})
	handler.AddTransport(transport.Options{})
	handler.AddTransport(transport.POST{})

	// Query cache
	handler.SetQueryCache(lru.New[*ast.QueryDocument](queryCacheLRUSize))

	// Enabling introspection
	handler.Use(extension.Introspection{})

	// APQ
	handler.Use(extension.AutomaticPersistedQuery{Cache: lru.New[string](automaticPersistedQueryCacheLRUSize)})

	// Complexity
	handler.Use(extension.FixedComplexityLimit(complexityLimit))

	// Unhandled errors logger
	handler.SetRecoverFunc(func(ctx context.Context, err interface{}) (userMessage error) {
		rh.Logger.LogrusLog.Errorf("unhandled error: %v", err)
		return gqlerror.Errorf("internal server error")
	})

	return handler
}
