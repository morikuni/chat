package usecase

import (
	"context"

	"github.com/morikuni/chat/chat/domain/model"
	"github.com/morikuni/chat/chat/domain/model/email"
	"github.com/morikuni/chat/chat/domain/model/password"
	"github.com/morikuni/chat/chat/domain/model/user"
	"github.com/morikuni/chat/eventsourcing"
	"github.com/pkg/errors"
)

func NewRegisterUser(store eventsourcing.EventStore) RegisterUser {
	return registerUser{
		store,
	}
}

type RegisterUser interface {
	Register(ctx context.Context, name, email, password string) (model.UserID, error)
}

type registerUser struct {
	store eventsourcing.EventStore
}

func (ru registerUser) Register(ctx context.Context, name, aEmail, aPassword string) (model.UserID, error) {
	n, err := user.ValidateName(name)
	if err != nil {
		return "", errors.WithStack(RaiseValidationError("name", err))
	}
	e, err := email.Validate(aEmail)
	if err != nil {
		return "", errors.WithStack(RaiseValidationError("email", err))
	}
	p, err := password.Validate(aPassword)
	if err != nil {
		return "", errors.WithStack(RaiseValidationError("password", err))
	}
	u := user.New(n, e, p)
	if err := ru.store.Save(ctx, u.Changes()); err != nil {
		return "", errors.WithMessage(err, "failed to save new user")
	}
	return u.ID(), nil
}
