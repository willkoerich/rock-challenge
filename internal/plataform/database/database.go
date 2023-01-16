package database

import (
	"context"
	"database/sql"
)

type (
	Driver interface {
		BeginTransaction(ctx context.Context) (Transaction, error)
		Exec(ctx context.Context, query string, args ...any) (ExecResult, error)
		ExecWithTransaction(ctx context.Context, transaction Transaction, query string, args ...any) (ExecResult, error)
		ExecuteInsert(ctx context.Context, command string, args ...interface{}) (int, error)
		ExecuteInsertWithTransaction(ctx context.Context, transaction Transaction, command string, args ...interface{}) (int, error)
		ExecuteQuerySingleElementCommand(ctx context.Context, command string, args ...interface{}) Result
		ExecuteQueryElementSetCommand(ctx context.Context, command string, args ...interface{}) (Results, error)
	}

	ExecResult interface {
		LastInsertId() (int64, error)
		RowsAffected() (int64, error)
	}

	Transaction interface {
		Rollback() error
		Commit() error
	}

	Result interface {
		Scan(dest ...interface{}) error
	}

	Results interface {
		Scan(dest ...interface{}) error
		Next() bool
		Close() error
	}

	PostgresDriver struct {
		db *sql.DB
	}

	PostgresTransaction struct {
		Tx *sql.Tx
	}
)

func (transaction PostgresTransaction) Commit() error {
	return transaction.Tx.Commit()
}

func (transaction PostgresTransaction) Rollback() error {
	return transaction.Tx.Rollback()
}

func NewPostgresDriver(db *sql.DB) Driver {
	return PostgresDriver{
		db: db,
	}
}

func (driver PostgresDriver) BeginTransaction(ctx context.Context) (Transaction, error) {
	tx, err := driver.db.BeginTx(ctx, nil)
	return PostgresTransaction{tx}, err
}

func (driver PostgresDriver) ExecuteInsert(ctx context.Context, command string, args ...interface{}) (int, error) {
	var id int
	err := driver.db.QueryRowContext(ctx, command, args...).Scan(&id)
	return id, err
}

func (driver PostgresDriver) ExecuteInsertWithTransaction(ctx context.Context, transaction Transaction, command string, args ...interface{}) (int, error) {
	var id int
	smt, err := driver.stmt(ctx, command, transaction)
	if err != nil {
		return 0, err
	}
	err = smt.QueryRowContext(ctx, args...).Scan(&id)
	return id, err
}

func (driver PostgresDriver) Exec(ctx context.Context, query string, args ...any) (ExecResult, error) {
	return driver.db.ExecContext(ctx, query, args...)
}

func (driver PostgresDriver) ExecWithTransaction(ctx context.Context, transaction Transaction, query string, args ...any) (ExecResult, error) {
	stmt, err := driver.stmt(ctx, query, transaction)
	if err != nil {
		return nil, err
	}
	return stmt.ExecContext(ctx, args...)
}

func (driver PostgresDriver) ExecuteQuerySingleElementCommand(ctx context.Context, command string, args ...interface{}) Result {
	return driver.db.QueryRowContext(ctx, command, args...)
}

func (driver PostgresDriver) ExecuteQueryElementSetCommand(ctx context.Context, command string, args ...interface{}) (Results, error) {
	return driver.db.QueryContext(ctx, command, args...)
}

func (driver PostgresDriver) stmt(ctx context.Context, command string, transaction Transaction) (stmt *sql.Stmt, err error) {

	postgresTransaction := transaction.(PostgresTransaction)
	if postgresTransaction.Tx != nil {
		stmt, err = postgresTransaction.Tx.PrepareContext(ctx, command)
	}

	stmt, err = driver.db.PrepareContext(ctx, command)
	return stmt, err
}
