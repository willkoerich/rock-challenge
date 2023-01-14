package database

import (
	"context"
	"database/sql"
)

type (
	Driver interface {
		ExecuteInsertCommand(ctx context.Context, command string, args ...interface{}) (int, error)
		Exec(ctx context.Context, query string, args ...any) (ExecResult, error)
		ExecuteQuerySingleElementCommand(ctx context.Context, command string, args ...interface{}) Result
		ExecuteQueryElementSetCommand(ctx context.Context, command string, args ...interface{}) (Results, error)
	}

	PostgresDriver struct {
		db *sql.DB
	}

	ExecResult interface {
		LastInsertId() (int64, error)
		RowsAffected() (int64, error)
	}

	Result interface {
		Scan(dest ...interface{}) error
	}

	Results interface {
		Scan(dest ...interface{}) error
		Next() bool
		Close() error
	}
)

func NewPostgresDriver(db *sql.DB) Driver {
	return PostgresDriver{
		db: db,
	}
}

func (driver PostgresDriver) ExecuteInsertCommand(ctx context.Context, command string, args ...interface{}) (int, error) {
	var id int
	err := driver.db.QueryRowContext(ctx, command, args...).Scan(&id)
	return id, err
}

func (driver PostgresDriver) Exec(ctx context.Context, query string, args ...any) (ExecResult, error) {
	return driver.db.ExecContext(ctx, query, args...)
}

func (driver PostgresDriver) ExecuteQuerySingleElementCommand(ctx context.Context, command string, args ...interface{}) Result {
	return driver.db.QueryRowContext(ctx, command, args...)
}

func (driver PostgresDriver) ExecuteQueryElementSetCommand(ctx context.Context, command string, args ...interface{}) (Results, error) {
	return driver.db.QueryContext(ctx, command, args...)
}
