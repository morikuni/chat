package infra

import (
	"context"
)

type TransactionManager interface {
	Exec(ctx context.Context, f func(context.Context) error) error
}
