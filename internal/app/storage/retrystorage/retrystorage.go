package retrystorage

import (
	"fmt"
	"time"

	"github.com/zhenyanesterkova/nepblog/internal/app/backoff"
	"github.com/zhenyanesterkova/nepblog/internal/app/config"
	"github.com/zhenyanesterkova/nepblog/internal/app/logger"
)

type Store interface {
	Ping() error
	Close() error
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
