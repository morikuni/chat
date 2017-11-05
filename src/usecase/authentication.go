package usecase

import (
	"context"

	"github.com/morikuni/chat/src/domain/event"
	"github.com/morikuni/chat/src/domain/model"
	"github.com/morikuni/chat/src/domain/model/aggregate"
	"github.com/morikuni/chat/src/domain/repository"
)

type Authentication interface {
	CreateAccount(ctx context.Context, email, password string) error
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

func (a authentication) CreateAccount(ctx context.Context, email, password string) error {
	em, verr := model.ValidateEmail(email)
	if verr != nil {
		return TranslateValidationError(verr, "email")
	}
	pw, verr := model.ValidatePassword(password)
	if verr != nil {
		return TranslateValidationError(verr, "password")
	}
	hash, err := pw.Hash()
	if err != nil {
		return err
	}
	id, err := a.accountRepository.GenerateID(ctx)
	if err != nil {
		return err
	}
	account, e := aggregate.NewAccount(id, em, hash)
	if err := a.accountRepository.Save(ctx, account); err != nil {
		return err
	}
	if err := a.eventPublisher.Publish(ctx, e); err != nil {
		return err
	}
	return nil
}
