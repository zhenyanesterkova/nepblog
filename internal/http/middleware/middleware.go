package middleware

import (
	"github.com/zhenyanesterkova/nepblog/internal/app/dataloaders"
	"github.com/zhenyanesterkova/nepblog/internal/app/logger"
	"github.com/zhenyanesterkova/nepblog/internal/feature/comment"
	"github.com/zhenyanesterkova/nepblog/internal/feature/post"
)

type MiddlewareStruct struct {
	Logger   logger.LogrusLogger
	respData *responseDataWriter
	loaders  *dataloaders.DataLoaders
}

func NewMiddlewareStruct(
	log logger.LogrusLogger,
	postLoader *post.LoaderByID,
	commentLoader *comment.LoaderByPostID,
) MiddlewareStruct {
	responseData := &responseData{
		status: 0,
		size:   0,
	}

	lw := responseDataWriter{
		responseData: responseData,
	}

	loaders := &dataloaders.DataLoaders{
		PostLoaderByID:        postLoader,
		CommentLoaderByPostID: commentLoader,
	}

	return MiddlewareStruct{
		Logger:   log,
		respData: &lw,
		loaders:  loaders,
	}
}
