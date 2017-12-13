package usecase

import (
	"context"

	"github.com/morikuni/chat/src/application"
	"github.com/morikuni/chat/src/domain"
	"github.com/morikuni/chat/src/domain/event"
	"github.com/morikuni/chat/src/domain/model"
	"github.com/morikuni/chat/src/domain/model/aggregate"
	"github.com/morikuni/chat/src/domain/repository"
	"github.com/morikuni/chat/src/infra"
)

type Authentication interface {
	CreateAccount(ctx context.Context, email, password string) error
	Login(ctx context.Context, email, password string) (model.UserID, error)
}

func NewAuthentication(
	accountRepository repository.Account,
	eventPublisher event.Publisher,
	transaction infra.TransactionManager,
) Authentication {
	return authentication{
		accountRepository,
		eventPublisher,
		transaction,
	}
}

type authentication struct {
	accountRepository repository.Account
	eventPublisher    event.Publisher
	transaction       infra.TransactionManager
}

func (a authentication) CreateAccount(ctx context.Context, email, password string) error {
	em, verr := model.ValidateEmail(email)
	if verr != nil {
		return application.TranslateValidationError(verr, "email")
	}
	pw, verr := model.ValidatePassword(password)
	if verr != nil {
		return application.TranslateValidationError(verr, "password")
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

	err = a.transaction.Exec(ctx, func(ctx context.Context) error {
		if err := a.accountRepository.Save(ctx, account); err != nil {
			return err
		}
		if err := a.eventPublisher.Publish(ctx, e); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		if _, ok := err.(domain.DuplicateEmailError); ok {
			return application.RaiseValidationError("email", "already in use")
		}
	}
	return nil
}

func (a authentication) Login(ctx context.Context, email, password string) (model.UserID, error) {
	em, verr := model.ValidateEmail(email)
	if verr != nil {
		return "", application.TranslateValidationError(verr, "email")
	}
	pw, verr := model.ValidatePassword(password)
	if verr != nil {
		return "", application.TranslateValidationError(verr, "password")
	}
	account, err := a.accountRepository.FindByEmail(ctx, em)
	if err != nil {
		if _, ok := err.(domain.NoSuchAggregateError); ok {
			return "", application.RaiseInvalidCredentialError()
		}
		return "", err
	}
	if err := account.LoginInfo.Password.Equal(pw); err != nil {
		if _, ok := err.(domain.PasswordMismatchError); ok {
			return "", application.RaiseInvalidCredentialError()
		}
		return "", err
	}

	return account.UserID, nil
}
