package infra

import (
	"context"

	"github.com/pkg/errors"
)

type Logger interface {
	Errorf(ctx context.Context, format string, args ...interface{})
}

type StackTraceError interface {
	error
	StackTrace() errors.StackTrace
}
