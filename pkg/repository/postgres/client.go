package postgres

import (
	"context"
	"fmt"
	"github.com/igilgyrg/gin-todo/pkg/utils"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Client interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

func NewClient(ctx context.Context, cfg *postgresConfig) (*pgxpool.Pool, error) {
	var pool *pgxpool.Pool
	var err error
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.username, cfg.password, cfg.host, cfg.port, cfg.database)
	err = utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()

		pool, err = pgxpool.New(ctx, dsn)
		if err != nil {
			return err
		}
		return nil
	}, cfg.maxAttempts, 5*time.Second)

	if err != nil {
		return nil, err
	}

	return pool, nil
}
