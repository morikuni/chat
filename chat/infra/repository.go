package infra

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type Repository interface {
	Append(ctx RepositoryContext, aType string, id string, data []byte) error
	Save(ctx RepositoryContext, aType string, id string, version int, data []byte) error
	Find(ctx RepositoryContext, aType string, id string) (version int, data []byte, err error)
}

type postgresRepo struct{}

func (r postgresRepo) Append(ctx RepositoryContext, aType string, id string, data []byte) error {
	_, err := ctx.ExecContext(ctx, `INSERT INTO aggregate (id, type, version, data) VALUES ($1, $2, $3, $4)`, id, aType, 0, data)
	if err != nil {
		if e, ok := err.(*pq.Error); ok && e.Code.Name() == "unique_violation" {
			return errors.WithStack(DuplicateError{fmt.Sprintf("duplicate id: type=%s id=%s data=%s", aType, id, string(data))})
		}
		return errors.Wrapf(err, "failed to insert: type=%s id=%s data=%s", aType, id, string(data))
	}
	return nil
}

func (r postgresRepo) Save(ctx RepositoryContext, aType string, id string, version int, data []byte) error {
	result, err := ctx.ExecContext(ctx, `UPDATE aggregate SET version = version + 1, data = $1 WHERE type = $2 AND id = $3 AND version = $4`, data, aType, id, version)
	if err != nil {
		return errors.Wrapf(err, "failed to update: type=%s id=%s data=%s", aType, id, string(data))
	}
	n, err := result.RowsAffected()
	if err != nil {
		return errors.Wrapf(err, "failed to get rows affected: type=%s id=%s data=%s", aType, id, string(data))
	} else if n == 0 {
		return NoRowsUpdatedError{fmt.Sprintf("no rows affected: type=%s id=%s data=%s", aType, id, string(data))}
	}
	return nil
}

func (r postgresRepo) Find(ctx RepositoryContext, aType string, id string) (int, []byte, error) {
	row := ctx.QueryRowContext(ctx, `SELECT version, data FROM aggregate WHERE type = $1 AND id = $2`, aType, id)
	var version int
	var data []byte
	err := row.Scan(&version, &data)
	if err == sql.ErrNoRows {
		return 0, nil, errors.WithStack(NoSuchRowError{})
	} else if err != nil {
		return 0, nil, errors.Wrapf(err, "failed to find: type=%s id=%s", aType, id)
	}
	return version, data, nil
}

type DuplicateError struct {
	message string
}

func (e DuplicateError) Error() string {
	return e.message
}

type NoSuchRowError struct{}

func (e NoSuchRowError) Error() string {
	return "no such row"
}

type NoRowsUpdatedError struct {
	message string
}

func (e NoRowsUpdatedError) Error() string {
	return e.message
}

type RepositoryContext interface {
	context.Context
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type executor interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type DBContext struct {
	context.Context
	db *sql.DB
}

func (c DBContext) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return c.db.ExecContext(ctx, query, args...)
}

func (c DBContext) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return c.db.QueryRowContext(ctx, query, args...)
}

type TxnContext struct {
	context.Context
	txn *sql.Tx
}

func (c TxnContext) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return c.txn.ExecContext(ctx, query, args...)
}

func (c TxnContext) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return c.txn.QueryRowContext(ctx, query, args...)
}

func (c TxnContext) Commit() error {
	return c.txn.Commit()
}

func (c TxnContext) Rollback() error {
	return c.txn.Rollback()
}
