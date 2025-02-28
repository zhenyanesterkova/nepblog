package comment

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type LoaderByPostIDRepository interface {
	FetchCommentsByPostID(ctx context.Context, postID []uuid.UUID) ([]Comment, error)
}

type LoaderByPostIDKey struct {
	PostID uuid.UUID
}

func NewConfiguredLoaderByPostID(repo LoaderByPostIDRepository, maxBatch int) *LoaderByPostID {
	return NewLoaderByPostID(LoaderByPostIDConfig{
		Wait:     2 * time.Millisecond,
		MaxBatch: maxBatch,
		Fetch: func(keys []LoaderByPostIDKey) ([][]Comment, []error) {
			items := make([][]Comment, len(keys))
			errors := make([]error, len(keys))

			postIDs := getUniquePostIDs(keys)

			comments, err := repo.FetchCommentsByPostID(
				context.Background(),
				postIDs,
			)
			if err != nil {
				for index := range keys {
					errors[index] = err
				}
			}

			groups := groupCommentsByPostID(comments)
			for i, key := range keys {
				if at, ok := groups[key.PostID]; ok {
					items[i] = at
				}
			}

			return items, errors
		},
	})
}

func getUniquePostIDs(keys []LoaderByPostIDKey) []uuid.UUID {
	mapping := make(map[uuid.UUID]bool)

	for _, key := range keys {
		mapping[key.PostID] = true
	}

	ids := make([]uuid.UUID, len(mapping))

	i := 0
	for key := range mapping {
		ids[i] = key
		i++
	}

	return ids
}

func groupCommentsByPostID(comments []Comment) map[uuid.UUID][]Comment {
	groups := make(map[uuid.UUID][]Comment)

	for _, at := range comments {
		groups[at.PostID] = append(groups[at.PostID], at)
	}

	return groups
}
