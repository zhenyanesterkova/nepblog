package middleware

import (
	"context"
	"net/http"

	"github.com/zhenyanesterkova/nepblog/internal/app/dataloaders"
)

func (lm MiddlewareStruct) DataLoaders() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(
				r.Context(),
				dataloaders.DataLoadersContextKey,
				lm.loaders,
			)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
