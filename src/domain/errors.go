package domain

import (
	"fmt"

	"github.com/morikuni/chat/src/infra"
	"github.com/pkg/errors"
)

type DomainError interface {
	infra.StackTraceError
	domainError()
}

type domainError struct {
	infra.StackTraceError
}

func (e domainError) domainError() {}

func ErrorOf(message string) DomainError {
	return domainError{errors.New(message).(infra.StackTraceError)}
}

type (
	ValidationError interface {
		DomainError
		validationError()
	}
	DuplicateEmailError struct {
		DomainError
	}
	NoSuchAggregateError struct {
		DomainError
	}
	PasswordMismatchError struct {
		DomainError
	}
)

type validationError struct {
	DomainError
}

func (ve validationError) validationError() {}

func RaiseValidationError(message string) ValidationError {
	return validationError{ErrorOf(message)}
}

func RaiseDuplicateEmailError(email string) error {
	return DuplicateEmailError{ErrorOf(fmt.Sprintf("email duplicated: %s", email))}
}

func RaiseNoSuchAggregateError() error {
	return NoSuchAggregateError{ErrorOf("no such aggregate")}
}

func RaisePasswordMismatchError() error {
	return PasswordMismatchError{ErrorOf("password mismatch")}
}
