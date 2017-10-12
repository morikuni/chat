package infra

import (
	"context"
)

type Logger interface {
	Errorf(ctx context.Context, format string, args ...interface{})
}
