package model

type Account struct {
	UserID    UserID
	LoginInfo LoginInfo
}

func NewAccount(id UserID, email Email, password PasswordHash) *Account {
	return &Account{
		id,
		LoginInfo{
			email,
			password,
		},
	}
}

type UserID int64

type LoginInfo struct {
	Email    Email
	Password PasswordHash
}
