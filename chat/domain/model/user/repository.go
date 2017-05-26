package user

import (
	"context"

	"github.com/morikuni/chat/chat/domain"
	"github.com/morikuni/chat/chat/domain/model"
	"github.com/morikuni/chat/eventsourcing"
	"github.com/pkg/errors"
)

func NewRepository(store eventsourcing.EventStore) model.UserRepository {
	return repository{
		store,
	}
}

type repository struct {
	store eventsourcing.EventStore
}

func (r repository) Find(ctx context.Context, id model.UserID) (model.User, error) {
	events, err := r.store.Find(ctx, string(id))
	switch {
	case err == eventsourcing.ErrNoEventsFound:
		return nil, errors.WithStack(domain.ErrNoSuchEntity)
	case err != nil:
		return nil, errors.Wrap(err, "failed to find events for user")
	}
	u := newUser()
	err = u.Replay(events)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to replay events")
	}
	return u, nil
}
