package database

import (
	"context"
	"database/sql"
)

type (
	Driver interface {
		ExecuteInsertCommand(ctx context.Context, command string, args ...interface{}) (int, error)
		ExecuteQuerySingleElementCommand(ctx context.Context, command string, args ...interface{}) Result
		ExecuteQueryElementSetCommand(ctx context.Context, command string, args ...interface{}) (Results, error)
	}

	PostgresDriver struct {
		db *sql.DB
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

func (driver PostgresDriver) ExecuteInsertCommand(ctx context.Context, command string, args ...interface{}) (int, error) {
	var id int
	err := driver.db.QueryRowContext(ctx, command, args...).Scan(&id)
	return id, err
}

func (driver PostgresDriver) ExecuteQuerySingleElementCommand(ctx context.Context, command string, args ...interface{}) Result {
	return driver.db.QueryRowContext(ctx, command, args...)
}

func (driver PostgresDriver) ExecuteQueryElementSetCommand(ctx context.Context, command string, args ...interface{}) (Results, error) {
	return driver.db.QueryContext(ctx, command, args...)
}
