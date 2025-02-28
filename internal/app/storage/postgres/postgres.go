package postgres

import (
	"context"
	"embed"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/zhenyanesterkova/nepblog/internal/app/logger"
	"github.com/zhenyanesterkova/nepblog/internal/feature/comment"
	"github.com/zhenyanesterkova/nepblog/internal/feature/post"
)

type PostgresStorage struct {
	pool *pgxpool.Pool
	log  logger.LogrusLogger
}

func New(dsn string, lg logger.LogrusLogger) (*PostgresStorage, error) {
	if err := runMigrations(dsn); err != nil {
		return nil, fmt.Errorf("failed to run DB migrations: %w", err)
	}
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create a connection pool: %w", err)
	}
	return &PostgresStorage{
		pool: pool,
		log:  lg,
	}, nil
}

//go:embed migrations/*.sql
var migrationsDir embed.FS

func runMigrations(dsn string) error {
	d, err := iofs.New(migrationsDir, "migrations")
	if err != nil {
		return fmt.Errorf("failed to return an iofs driver: %w", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, dsn)
	if err != nil {
		return fmt.Errorf("failed to get a new migrate instance: %w", err)
	}
	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("failed to apply migrations to the DB: %w", err)
		}
	}
	return nil
}

func (psg *PostgresStorage) FetchPosts(ctx context.Context, ids []uuid.UUID) ([]post.Post, error) {
	return []post.Post{}, nil
}

func (psg *PostgresStorage) FetchCommentsByPostID(ctx context.Context, postID []uuid.UUID) ([]comment.Comment, error) {
	return []comment.Comment{}, nil
}

func (psg *PostgresStorage) Ping() error {
	if err := psg.pool.Ping(context.TODO()); err != nil {
		return fmt.Errorf("failed to ping the DB: %w", err)
	}

	return nil
}

func (psg *PostgresStorage) Close() error {
	psg.pool.Close()
	return nil
}
