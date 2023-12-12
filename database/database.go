package database

import (
	"context"
	"database/sql"
)

type Query interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type Transaction interface {
	Commit() error
	Rollback() error
	Query
}

type Database interface {
	PingContext(ctx context.Context) error
	Close() error
	BeginTx(ctx context.Context, opts *sql.TxOptions) (Transaction, error)
	Query
}

type DB struct {
	conn *sql.DB
}

func NewDB(db *sql.DB) *DB {
	return &DB{conn: db}
}

func (db DB) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return db.conn.ExecContext(ctx, query, args...)
}

func (db DB) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return db.conn.PrepareContext(ctx, query)
}

func (db DB) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return db.conn.QueryContext(ctx, query, args...)
}

func (db DB) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return db.conn.QueryRowContext(ctx, query, args...)
}

func (db DB) PingContext(ctx context.Context) error {
	return db.conn.PingContext(ctx)
}

func (db DB) Close() error {
	return db.conn.Close()
}

func (db DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (Transaction, error) {
	return db.conn.BeginTx(ctx, opts)
}
