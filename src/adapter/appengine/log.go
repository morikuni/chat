package appengine

import (
	"context"

	"github.com/morikuni/chat/src/infra"
	"google.golang.org/appengine/log"
)

func NewLogger() infra.Logger {
	return logger{}
}

type logger struct{}

func (l logger) Errorf(ctx context.Context, format string, args ...interface{}) {
	log.Errorf(ctx, format, args...)
}

func (l logger) Debugf(ctx context.Context, format string, args ...interface{}) {
	log.Debugf(ctx, format, args...)
}
