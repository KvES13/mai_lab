package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type Client interface {
	Begin(context.Context) (pgx.Tx, error)
	BeginFunc(ctx context.Context, f func(pgx.Tx) error) error
	BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) error
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

type pgConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

// NewPgConfig creates new pg config instance
func NewPgConfig(username string, password string, host string, port string, database string) *pgConfig {
	return &pgConfig{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
		Database: database,
	}
}

func NewClient(ctx context.Context, cfg *pgConfig) (pool *pgxpool.Pool, err error) {

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	pool, err = pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
		return nil, err
	}

	return pool, nil
}
