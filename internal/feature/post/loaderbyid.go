package post

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type LoaderByIDRepository interface {
	FetchPosts(ctx context.Context, ids []uuid.UUID) ([]Post, error)
}

type LoaderByIDKey struct {
	ID uuid.UUID
}

func NewConfiguredLoaderByID(repo LoaderByIDRepository, maxBatch int) *LoaderByID {
	return NewLoaderByID(LoaderByIDConfig{
		Wait:     2 * time.Millisecond,
		MaxBatch: maxBatch,
		Fetch: func(keys []LoaderByIDKey) ([]*Post, []error) {
			items := make([]*Post, len(keys))
			errors := make([]error, len(keys))

			ctx, cancel := context.WithTimeout(
				context.Background(),
				50*time.Millisecond*time.Duration(len(keys)),
			)
			defer cancel()

			ids := getUniquePostIDs(keys)

			posts, err := repo.FetchPosts(
				ctx,
				ids,
			)
			if err != nil {
				for i := range keys {
					errors[i] = err
				}
			}

			groups := groupPostsByID(posts)
			for i, key := range keys {
				if a, ok := groups[key.ID]; ok {
					items[i] = &a
				}
			}

			return items, errors
		},
	})
}

func getUniquePostIDs(keys []LoaderByIDKey) []uuid.UUID {
	mapping := make(map[uuid.UUID]bool)

	for _, key := range keys {
		mapping[key.ID] = true
	}

	ids := make([]uuid.UUID, len(mapping))

	i := 0
	for key := range mapping {
		ids[i] = key
		i++
	}

	return ids
}

func groupPostsByID(posts []Post) map[uuid.UUID]Post {
	groups := make(map[uuid.UUID]Post)

	for _, a := range posts {
		groups[a.ID] = a
	}

	return groups
}
