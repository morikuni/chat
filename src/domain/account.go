package domain

import (
	"regexp"

	"github.com/google/uuid"
	"github.com/morikuni/chat/src/errors"
	"github.com/morikuni/failure"
)

type Account struct {
	ID   AccountID
	Name AccountName
}

func NewAccount(name AccountName) *Account {
	return &Account{
		generateAccountID(),
		name,
	}
}

type AccountID string

func NewAccountID(id string) (AccountID, error) {
	if id == "" {
		return "", failure.New(errors.InvalidArgument,
			failure.Message("account id must not be empty"),
		)
	}

	return AccountID(id), nil
}

func generateAccountID() AccountID {
	return AccountID(uuid.New().String())
}

type AccountName string

var accountNameReg = regexp.MustCompile("[a-z0-9A-Z_-]+")

func NewAccountName(name string) (AccountName, error) {
	if !accountNameReg.MatchString(name) {
		return "", failure.New(errors.InvalidArgument, failure.Message("name contains invalid characters"))
	}

	if l := len(name); l == 0 || l > 20 {
		return "", failure.New(errors.InvalidArgument, failure.Message("name must be from 1 to 20 characters"))
	}

	return AccountName(name), nil
}
