package impl

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
