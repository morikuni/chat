package datastore

import (
	"context"

	"github.com/morikuni/chat/src/domain"
	"github.com/morikuni/chat/src/domain/model"
	"github.com/morikuni/chat/src/domain/model/aggregate"
	"github.com/morikuni/chat/src/domain/repository"
	"github.com/morikuni/chat/src/infra"
	"github.com/pkg/errors"
	"google.golang.org/appengine/datastore"
)

const (
	AccountKind = "Account"
	EmailKind   = "Email"
)

func NewAccountRepository(transaction infra.TransactionManager) repository.Account {
	return account{
		transaction,
	}
}

type account struct {
	transaction infra.TransactionManager
}

type email struct {
	userID model.UserID
}

func (account) GenerateID(ctx context.Context) (model.UserID, error) {
	l, _, err := datastore.AllocateIDs(ctx, AccountKind, nil, 1)
	if err != nil {
		return 0, errors.Wrap(err, "failed to generate user ID")
	}
	return model.UserID(l), nil
}

func (a account) Save(ctx context.Context, account *aggregate.Account) error {
	accountKey := datastore.NewKey(ctx, AccountKind, "", int64(account.UserID), nil)
	emailKey := datastore.NewKey(ctx, EmailKind, string(account.LoginInfo.Email), 0, nil)
	return a.transaction.Exec(ctx, func(ctx context.Context) error {
		shouldSaveEmail := true
		var em email
		err := datastore.Get(ctx, emailKey, &em)
		if err != nil && err != datastore.ErrNoSuchEntity {
			return errors.Wrap(err, "failed to get email")
		}
		if err == nil {
			if em.userID != account.UserID {
				return domain.RaiseDuplicateEmailError(string(account.LoginInfo.Email))
			}
			shouldSaveEmail = false
		}

		keys := []*datastore.Key{accountKey}
		values := []interface{}{account}
		if shouldSaveEmail {
			keys = append(keys, emailKey)
			values = append(values, &em)
		}

		em = email{account.UserID}
		if _, err = datastore.PutMulti(ctx, keys, values); err != nil {
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
