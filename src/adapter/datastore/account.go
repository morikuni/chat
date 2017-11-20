package datastore

import (
	"context"
	"strconv"

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

type emailEntity struct {
	UserID model.UserID
}

func (account) GenerateID(ctx context.Context) (model.UserID, error) {
	l, _, err := datastore.AllocateIDs(ctx, AccountKind, nil, 1)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate user ID")
	}
	return model.UserID(strconv.FormatInt(l, 10)), nil
}

func (a account) Save(ctx context.Context, account *aggregate.Account) error {
	accountKey := datastore.NewKey(ctx, AccountKind, string(account.UserID), 0, nil)
	emailKey := datastore.NewKey(ctx, EmailKind, string(account.LoginInfo.Email), 0, nil)
	return a.transaction.Exec(ctx, func(ctx context.Context) error {
		shouldSaveEmail := true
		var em emailEntity
		err := datastore.Get(ctx, emailKey, &em)
		if err != nil && err != datastore.ErrNoSuchEntity {
			return errors.Wrap(err, "failed to get email")
		}
		if err == nil {
			if em.UserID != account.UserID {
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

		em = emailEntity{account.UserID}
		if _, err = datastore.PutMulti(ctx, keys, values); err != nil {
			return errors.Wrap(err, "failed to save account")
		}
		return nil
	})
}

func (account) Find(ctx context.Context, id model.UserID) (*aggregate.Account, error) {
	key := datastore.NewKey(ctx, AccountKind, string(id), 0, nil)
	var account aggregate.Account
	if err := datastore.Get(ctx, key, &account); err != nil {
		if err == datastore.ErrNoSuchEntity {
			return nil, domain.RaiseNoSuchAggregateError()
		}
		return nil, errors.Wrap(err, "failed to get account")
	}
	return &account, nil
}
