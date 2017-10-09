package infra

import (
	"context"
	"google.golang.org/appengine/log"
)

func NewLogger() Logger {
	return logger{}
}

type Logger interface {
	Errorf(ctx context.Context, format string, args ...interface{})
}

type logger struct{}

func (l logger) Errorf(ctx context.Context, format string, args ...interface{}) {
	log.Errorf(ctx, format, args...)
}
