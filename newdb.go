package mydb

import (
	"time"
	"context"
	"database/sql"
)

type Store interface{
	Begin() (*sql.Tx, error)
    BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
    Close() error
    Exec(query string, args ...interface{}) (sql.Result, error)
    ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
    Ping() error
    PingContext(ctx context.Context) error
    Prepare(query string) (*sql.Stmt, error)
    PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
    Query(query string, args ...interface{}) (*sql.Rows, error)
    QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
    QueryRow(query string, args ...interface{}) *sql.Row
    QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
   // SetConnMaxIdleTime(d time.Duration)
    SetConnMaxLifetime(d time.Duration)
    SetMaxIdleConns(n int)
    SetMaxOpenConns(n int)
}