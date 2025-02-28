package retrystorage

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/zhenyanesterkova/nepblog/internal/app/backoff"
	"github.com/zhenyanesterkova/nepblog/internal/app/config"
	"github.com/zhenyanesterkova/nepblog/internal/app/logger"
	"github.com/zhenyanesterkova/nepblog/internal/feature/comment"
	"github.com/zhenyanesterkova/nepblog/internal/feature/post"
)

type Store interface {
	Ping() error
	Close() error
	FetchPosts(ctx context.Context, ids []uuid.UUID) ([]post.Post, error)
	FetchCommentsByPostID(ctx context.Context, postID []uuid.UUID) ([]comment.Comment, error)
}

type RetryStorage struct {
	storage    Store
	backoff    *backoff.Backoff
	logger     logger.LogrusLogger
	checkRetry func(error) bool
}

func New(
	cfg config.DataBaseConfig,
	loggerInst logger.LogrusLogger,
	bf *backoff.Backoff,
	checkRetryFunc func(error) bool,
	store Store,
) (
	*RetryStorage,
	error,
) {
	retryStore := &RetryStorage{
		checkRetry: checkRetryFunc,
		backoff:    bf,
		logger:     loggerInst,
		storage:    store,
	}

	return retryStore, nil
}

func (rs *RetryStorage) Ping() error {
	err := rs.storage.Ping()
	if rs.checkRetry(err) {
		err = rs.retry(func() error {
			err = rs.storage.Ping()
			if err != nil {
				return fmt.Errorf("failed retry ping: %w", err)
			}
			return nil
		})
	}
	if err != nil {
		return fmt.Errorf("failed ping: %w", err)
	}
	return nil
}

func (rs *RetryStorage) FetchPosts(ctx context.Context, ids []uuid.UUID) ([]post.Post, error) {
	posts, err := rs.storage.FetchPosts(ctx, ids)
	if rs.checkRetry(err) {
		err = rs.retry(func() error {
			posts, err = rs.storage.FetchPosts(ctx, ids)
			if err != nil {
				return fmt.Errorf("failed retry FetchPosts: %w", err)
			}
			return nil
		})
	}
	if err != nil {
		return posts, fmt.Errorf("failed FetchPosts: %w", err)
	}
	return posts, nil
}

func (rs *RetryStorage) FetchCommentsByPostID(ctx context.Context, postID []uuid.UUID) ([]comment.Comment, error) {
	posts, err := rs.storage.FetchCommentsByPostID(ctx, postID)
	if rs.checkRetry(err) {
		err = rs.retry(func() error {
			posts, err = rs.storage.FetchCommentsByPostID(ctx, postID)
			if err != nil {
				return fmt.Errorf("failed retry FetchCommentsByPostID: %w", err)
			}
			return nil
		})
	}
	if err != nil {
		return posts, fmt.Errorf("failed FetchCommentsByPostID: %w", err)
	}
	return posts, nil
}

func (rs *RetryStorage) Close() error {
	if err := rs.storage.Close(); err != nil {
		return fmt.Errorf("failed close DB: %w", err)
	}
	return nil
}

func (rs *RetryStorage) retry(work func() error) error {
	log := rs.logger.LogrusLog
	defer rs.backoff.Reset()
	for {
		log.Debug("attempt to repeat ...")
		err := work()

		if err == nil {
			return nil
		}

		if rs.checkRetry(err) {
			var delay time.Duration
			if delay = rs.backoff.Next(); delay == backoff.Stop {
				return err
			}
			time.Sleep(delay)
		}
	}
}
