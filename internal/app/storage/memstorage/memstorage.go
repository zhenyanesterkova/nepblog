package memstorage

import (
	"github.com/zhenyanesterkova/nepblog/internal/feature/comment"
	"github.com/zhenyanesterkova/nepblog/internal/feature/post"
)

type MemStorage struct {
	posts    []*post.Post
	comments []*comment.Comment
}

func New() *MemStorage {
	return &MemStorage{
		posts:    make([]*post.Post, 0, 10),
		comments: make([]*comment.Comment, 0, 10),
	}
}

func (s *MemStorage) Close() error {
	return nil
}

func (s *MemStorage) Ping() error {
	return nil
}
