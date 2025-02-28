package storage

import (
	"errors"
	"fmt"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/zhenyanesterkova/nepblog/internal/app/backoff"
	"github.com/zhenyanesterkova/nepblog/internal/app/config"
	"github.com/zhenyanesterkova/nepblog/internal/app/logger"
	"github.com/zhenyanesterkova/nepblog/internal/app/storage/memstorage"
	"github.com/zhenyanesterkova/nepblog/internal/app/storage/postgres"
	"github.com/zhenyanesterkova/nepblog/internal/app/storage/retrystorage"
)

type Store interface {
	Ping() error
	Close() error
}

func NewStore(conf *config.Config, log logger.LogrusLogger) (Store, error) {
	if conf.DBConfig.PostgresConfig != nil {
		log.LogrusLog.Debugln("create postgres storage")
		store, err := postgres.New(conf.DBConfig.PostgresConfig.DSN, log)
		if err != nil {
			return nil, fmt.Errorf("failed create postgres storage: %w", err)
		}

		backoffInst := backoff.New(
			conf.RetryConfig.MinDelay,
			conf.RetryConfig.MaxDelay,
			conf.RetryConfig.MaxAttempt,
		)

		checkRetryFunc := func(err error) bool {
			var pgErr *pgconn.PgError
			var pgErrConn *pgconn.ConnectError
			res := false
			if errors.As(err, &pgErr) {
				res = pgerrcode.IsConnectionException(pgErr.Code)
			} else if errors.As(err, &pgErrConn) {
				res = true
			}
			return res
		}

		retryStore, err := retrystorage.New(conf.DBConfig, log, backoffInst, checkRetryFunc, store)
		if err != nil {
			log.LogrusLog.Errorf("failed create storage: %v", err)
			return nil, fmt.Errorf("failed create storage: %w", err)
		}

		return retryStore, nil
	}

	log.LogrusLog.Debugln("create memory storage")
	store := memstorage.New()

	return store, nil
}
