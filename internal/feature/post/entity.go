package post

import (
	"time"

	"github.com/google/uuid"

	"github.com/zhenyanesterkova/nepblog/internal/feature/comment"
)

type Post struct {
	CreatedAt       time.Time
	ID              uuid.UUID
	UserID          uuid.UUID
	Title           string
	Content         string
	AllowedComments bool
	Comments        []comment.Comment
}
