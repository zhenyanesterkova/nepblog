package post

import (
	"github.com/zhenyanesterkova/nepblog/internal/feature/comment"
	"github.com/zhenyanesterkova/nepblog/internal/gql/model"
)

func MapOneToGqlModel(post Post) *model.Post {
	comments := &model.CommentList{
		Items: comment.MapManyToGqlModels(post.Comments),
	}
	return &model.Post{
		CreatedAt:       post.CreatedAt,
		ID:              post.ID,
		UserID:          post.UserID,
		Content:         post.Content,
		Title:           post.Title,
		AllowedComments: post.AllowedComments,
		Comments:        comments,
	}
}

func MapManyToGqlModels(posts []Post) []*model.Post {
	items := make([]*model.Post, len(posts))

	for i, entity := range posts {
		items[i] = MapOneToGqlModel(entity)
	}

	return items
}
