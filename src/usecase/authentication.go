package usecase

import (
	"context"

	"github.com/morikuni/chat/src/domain/event"
	"github.com/morikuni/chat/src/domain/model"
	"github.com/morikuni/chat/src/domain/model/aggregate"
	"github.com/morikuni/chat/src/domain/repository"
)

type Authentication interface {
	CreateAccount(ctx context.Context, email, password string) (model.UserID, error)
}

func NewAuthentication(
	accountRepository repository.Account,
	eventPublisher event.Publisher,
) Authentication {
	return authentication{
		accountRepository,
		eventPublisher,
	}
}

type authentication struct {
	accountRepository repository.Account
	eventPublisher    event.Publisher
}

func (a authentication) CreateAccount(ctx context.Context, email, password string) (model.UserID, error) {
	em, verr := model.ValidateEmail(email)
	if verr != nil {
		return 0, TranslateValidationError(verr, "email")
	}
	pw, verr := model.ValidatePassword(password)
	if verr != nil {
		return 0, TranslateValidationError(verr, "password")
	}
	hash, err := pw.Hash()
	if err != nil {
		return 0, err
	}
	id, err := a.accountRepository.GenerateID(ctx)
	if err != nil {
		return 0, err
	}
	account, e := aggregate.NewAccount(id, em, hash)
	if err := a.accountRepository.Save(ctx, account); err != nil {
		return 0, err
	}
	if err := a.eventPublisher.Publish(c, e); err != nil {
		return 0, err
	}
	return account.UserID, nil
}
