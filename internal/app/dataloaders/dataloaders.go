package dataloaders

import (
	"github.com/zhenyanesterkova/nepblog/internal/feature/comment"
	"github.com/zhenyanesterkova/nepblog/internal/feature/post"
)

type ContextKey int

const (
	DataLoadersContextKey ContextKey = iota
)

type DataLoaders struct {
	PostLoaderByID        *post.LoaderByID
	CommentLoaderByPostID *comment.LoaderByPostID
}
