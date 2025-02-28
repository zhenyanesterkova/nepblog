package post

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	CreatedAt       time.Time
	ID              uuid.UUID
	UserID          uuid.UUID
	Title           string
	Content         string
	AllowedComments bool
}
