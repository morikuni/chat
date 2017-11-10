package datastore

import (
	"context"

	"github.com/morikuni/chat/src/domain/model"
	"github.com/morikuni/chat/src/domain/model/aggregate"
	"github.com/morikuni/chat/src/domain/repository"
	"github.com/morikuni/chat/src/infra"
	"github.com/pkg/errors"
	"google.golang.org/appengine/datastore"
)

const (
	AccountKind = "Account"
)

func NewAccountRepository(transaction infra.TransactionManager) repository.Account {
	return account{
		transaction,
	}
}

type account struct {
	transaction infra.TransactionManager
}

func (account) GenerateID(ctx context.Context) (model.UserID, error) {
	l, _, err := datastore.AllocateIDs(ctx, AccountKind, nil, 1)
	if err != nil {
		return 0, errors.Wrap(err, "failed to generate user ID")
	}
	return model.UserID(l), nil
}

func (a account) Save(ctx context.Context, account *aggregate.Account) error {
	key := datastore.NewKey(ctx, AccountKind, "", int64(account.UserID), nil)
	return a.transaction.Exec(ctx, func(ctx context.Context) error {
		_, err := datastore.Put(ctx, key, account)
		if err != nil {
			return errors.Wrap(err, "failed to save account")
		}
		return nil
	})
}

func (account) Find(ctx context.Context, id model.UserID) (*aggregate.Account, error) {
	key := datastore.NewKey(ctx, AccountKind, "", int64(id), nil)
	var account aggregate.Account
	if err := datastore.Get(ctx, key, &account); err != nil {
		return nil, errors.Wrap(err, "failed to find account")
	}
	return &account, nil
}
