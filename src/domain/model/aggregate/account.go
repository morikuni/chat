package aggregate

import (
	"github.com/morikuni/chat/src/domain/model"
)

type Account struct {
	UserID    model.UserID
	LoginInfo LoginInfo
}

func NewAccount(id model.UserID, email model.Email, password model.PasswordHash) *Account {
	return &Account{
		id,
		LoginInfo{
			email,
			password,
		},
	}
}

type LoginInfo struct {
	Email    model.Email
	Password model.PasswordHash
}
