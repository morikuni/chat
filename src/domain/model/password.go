package model

import (
	"github.com/morikuni/chat/src/domain"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

const (
	Cost = 10
)

type Password []byte

func ValidatePassword(password string) (Password, domain.ValidationError) {
	if len(password) == 0 {
		return nil, domain.RaiseValidationError("cannot be empty")
	}
	return Password(password), nil
}

func (p Password) Hash() (PasswordHash, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(p), Cost)
	if err != nil {
		return PasswordHash{}, errors.Wrap(err, "failed to grenerate hash")
	}
	return PasswordHash(h), nil
}

type PasswordHash []byte

func (p PasswordHash) Equal(password Password) error {
	return errors.Wrap(bcrypt.CompareHashAndPassword([]byte(p), []byte(password)), "failed to compare password")
}
