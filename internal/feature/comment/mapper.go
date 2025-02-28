package comment

import "github.com/zhenyanesterkova/nepblog/internal/gql/model"

func MapOneToGqlModel(comment Comment) *model.Comment {
	childComments := &model.CommentList{
		Items: MapManyToGqlModels(comment.ChildComments),
	}
	return &model.Comment{
		PostID:        comment.PostID,
		CreatedAt:     comment.CreatedAt,
		ID:            comment.ID,
		UserID:        comment.UserID,
		Data:          comment.Data,
		ChildComments: childComments,
		ParentID:      comment.ParentComment,
	}
}

func MapManyToGqlModels(comments []Comment) []*model.Comment {
	items := make([]*model.Comment, len(comments))

	for i, entity := range comments {
		items[i] = MapOneToGqlModel(entity)
	}

	return items
}
