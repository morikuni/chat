package usecase

import (
	"context"

	"github.com/morikuni/chat/chat/domain"
	"github.com/morikuni/chat/chat/domain/model"
	"github.com/morikuni/chat/chat/domain/model/email"
	"github.com/morikuni/chat/chat/domain/model/password"
	"github.com/morikuni/chat/chat/domain/model/user"
	"github.com/pkg/errors"
)

func NewRegisterUser(userRepo model.UserRepository) RegisterUser {
	return registerUser{
		userRepo,
	}
}

type RegisterUser interface {
	Register(ctx context.Context, name, email, password string) (model.UserID, error)
}

type registerUser struct {
	userRepo model.UserRepository
}

func (ru registerUser) Register(ctx context.Context, name, aEmail, aPassword string) (model.UserID, error) {
	n, err := user.ValidateName(name)
	if err != nil {
		switch e := errors.Cause(err).(type) {
		case domain.ValidationError:
			return "", errors.WithStack(RaiseValidationError("name", e))
		default:
			return "", errors.WithMessage(err, "failed to create name")
		}
	}
	e, err := email.Validate(aEmail)
	if err != nil {
		switch e := errors.Cause(err).(type) {
		case domain.ValidationError:
			return "", errors.WithStack(RaiseValidationError("email", e))
		default:
			return "", errors.WithMessage(err, "failed to create email")
		}
	}
	p, err := password.Validate(aPassword)
	if err != nil {
		switch e := errors.Cause(err).(type) {
		case domain.ValidationError:
			return "", errors.WithStack(RaiseValidationError("password", e))
		default:
			return "", errors.WithMessage(err, "failed to create password")
		}
	}
	u := user.New(n, e, p)
	err = ru.userRepo.Append(ctx, u)
	if err != nil {
		return "", errors.WithMessage(err, "failed to save new user")
	}
	return u.ID(), nil
}
