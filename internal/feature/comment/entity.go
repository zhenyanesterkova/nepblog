package comment

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	CreatedAt time.Time
	ID        uuid.UUID
	Data      string
	UserID    uuid.UUID
	PostID    uuid.UUID
}
