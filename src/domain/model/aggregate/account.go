package aggregate

import (
	"github.com/morikuni/chat/src/domain/event"
	"github.com/morikuni/chat/src/domain/model"
)

type Account struct {
	UserID    model.UserID
	LoginInfo LoginInfo
}

func NewAccount(id model.UserID, email model.Email, password model.PasswordHash) (*Account, event.Event) {
	return &Account{
			id,
			LoginInfo{
				email,
				password,
			},
		}, event.AccountCreated{
			id,
			email,
		}
}

type LoginInfo struct {
	Email    model.Email
	Password model.PasswordHash
}
