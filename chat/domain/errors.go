package domain

import (
	"fmt"
	"github.com/morikuni/chat/eventsourcing"
)

type DomainError interface {
	error
	domainError()
}

type domainError struct {
	message string
}

func (e domainError) Error() string {
	return e.message
}

func (e domainError) domainError() {}

func ErrorOf(message string) DomainError {
	return domainError{message}
}

type (
	NoSuchEntityError       struct{ DomainError }
	EntityAlreadyExistError struct{ DomainError }
	UnexpectedEventError    struct{ DomainError }
)

var (
	ErrNoSuchEntity       = NoSuchEntityError{ErrorOf("no such entity")}
	ErrEntityAlreadyExist = EntityAlreadyExistError{ErrorOf("entity already exist")}
)

func RaiseValidationError(message string) ValidationError {
	return validationError{ErrorOf(message)}
}

func RaiseUnexpectedEventError(event eventsourcing.Event) UnexpectedEventError {
	return UnexpectedEventError{ErrorOf(fmt.Sprintf("unexpected event: %#v", event))}
}

type ValidationError interface {
	DomainError
	validationError()
}

type validationError struct {
	DomainError
}

func (e validationError) validationError() {}
