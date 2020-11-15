package usecase

import (
	"context"

	"github.com/morikuni/chat/src/domain"
	"github.com/morikuni/failure"
)

type account struct {
	accountRepo AccountRepository
}

var _ Account = (*account)(nil)

func NewAccount(
	accountRepo AccountRepository,
) Account {
	return &account{
		accountRepo,
	}
}

type CreateAccountRequest struct {
	Name domain.AccountName
}

type CreateAccountResponse struct {
	Account *domain.Account
}

func (a *account) CreateAccount(ctx context.Context, request *CreateAccountRequest) (*CreateAccountResponse, error) {
	account := domain.NewAccount(request.Name)

	err := a.accountRepo.Save(ctx, account)
	if err != nil {
		return nil, failure.Wrap(err)
	}

	return &CreateAccountResponse{account}, nil
}
