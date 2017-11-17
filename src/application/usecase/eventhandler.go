package usecase

import (
	"context"

	"github.com/morikuni/chat/src/domain/event"
	"github.com/morikuni/chat/src/infra"
)

func NewEventHandler(logger infra.Logger) EventHandler {
	return eventHandler{logger}
}

type EventHandler interface {
	Handle(c context.Context, event event.Event) error
}

type eventHandler struct {
	logger infra.Logger
}

func (eh eventHandler) Handle(ctx context.Context, e event.Event) error {
	switch e.(type) {
	case event.AccountCreated:
		eh.logger.Debugf(ctx, "event received: %#v", e)
	default:
		return RaiseUnknownEventError(e)
	}
	return nil
}
