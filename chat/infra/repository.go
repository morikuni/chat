package infra

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type Repository interface {
	Append(ctx context.Context, aType string, id string, object []byte) error
	Save(ctx context.Context, aType string, id string, version int, object []byte) error
	Find(ctx context.Context, aType string, id string) ([]byte, error)
}

type postgresRepo struct {
	db *sql.DB
}

func (r postgresRepo) Append(ctx context.Context, aType string, id string, object []byte) error {
	_, err := r.db.ExecContext(ctx, `INSERT INTO aggregate (id, type, version, object) VALUES ($1, $2, $3, $4)`, id, aType, 0, object)
	if err != nil {
		if e, ok := err.(*pq.Error); ok && e.Code.Name() == "unique_violation" {
			return errors.WithStack(DuplicateError{fmt.Sprintf("duplicate id: type=%s id=%s object=%s", aType, id, string(object))})
		}
		return errors.Wrapf(err, "failed to insert: type=%s id=%s object=%s", aType, id, string(object))
	}
	return nil
}

func (r postgresRepo) Save(ctx context.Context, aType string, id string, version int, object []byte) error {
	result, err := r.db.ExecContext(ctx, `UPDATE aggregate SET version = version + 1, object = $1 WHERE type = $2 AND id = $3 AND version = $4`, object, aType, id, version)
	if err != nil {
		return errors.Wrapf(err, "failed to update: type=%s id=%s object=%s", aType, id, string(object))
	}
	n, err := result.RowsAffected()
	if err != nil {
		return errors.Wrapf(err, "failed to get rows affected: type=%s id=%s object=%s", aType, id, string(object))
	} else if n == 0 {
		return NoRowsUpdatedError{fmt.Sprintf("no rows affected: type=%s id=%s object=%s", aType, id, string(object))}
	}
	return nil
}

func (r postgresRepo) Find(ctx context.Context, aType string, id string) (int, []byte, error) {
	row := r.db.QueryRowContext(ctx, `SELECT version, object FROM aggregate WHERE type = $1 AND id = $2`, aType, id)
	var version int
	var object []byte
	err := row.Scan(&version, &object)
	if err == sql.ErrNoRows {
		return 0, nil, errors.WithStack(NoSuchRowError{})
	} else if err != nil {
		return 0, nil, errors.Wrapf(err, "failed to find: type=%s id=%s")
	}
	return version, object, nil
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
