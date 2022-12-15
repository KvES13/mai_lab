package storage

import (
	"context"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgconn"
)

type PostgreSQLClient interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	//BeginFunc(ctx context.Context, f func(pgx.Tx) error) error
	//BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) error
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}
