package memstorage

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/zhenyanesterkova/nepblog/internal/feature/comment"
	"github.com/zhenyanesterkova/nepblog/internal/feature/post"
)

type MemStorage struct {
	posts    map[uuid.UUID]*post.Post
	comments map[uuid.UUID][]*comment.Comment
}

func New() *MemStorage {
	store := &MemStorage{
		posts:    make(map[uuid.UUID]*post.Post),
		comments: make(map[uuid.UUID][]*comment.Comment),
	}
	uuid, _ := uuid.NewUUID()
	store.posts[uuid] = &post.Post{
		CreatedAt:       time.Now(),
		ID:              uuid,
		UserID:          uuid,
		Title:           "Post #1",
		Content:         `<h1>Hello from Post #1</h1>`,
		AllowedComments: true,
	}
	return store
}

func (s *MemStorage) FetchPosts(ctx context.Context, ids []uuid.UUID) ([]post.Post, error) {
	res := make([]post.Post, 0, len(ids))
	for _, v := range ids {
		res = append(res, *s.posts[v])
	}

	return res, nil
}

func (s *MemStorage) FetchCommentsByPostID(ctx context.Context, postID []uuid.UUID) ([]comment.Comment, error) {
	res := []comment.Comment{}
	for _, v := range postID {
		for _, comment := range s.comments[v] {
			res = append(res, *comment)
		}
	}

	return res, nil
}

func (s *MemStorage) Close() error {
	return nil
}

func (s *MemStorage) Ping() error {
	return nil
}
