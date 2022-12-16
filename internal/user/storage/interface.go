package storage

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type PostgreSQLClient interface {
	//Begin(ctx context.Context) (pgx.Tx, error)
	//BeginFunc(ctx context.Context, f func(pgx.Tx) error) error
	//	BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) error
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	//Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
}