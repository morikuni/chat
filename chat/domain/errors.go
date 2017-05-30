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
	ValidationError         struct{ DomainError }
	UnexpectedCommandError  struct{ DomainError }
	UnexpectedEventError    struct{ DomainError }
)

var (
	ErrNoSuchEntity       = NoSuchEntityError{ErrorOf("no such entity")}
	ErrEntityAlreadyExist = EntityAlreadyExistError{ErrorOf("entity already exist")}
)

func RaiseValidationError(message string) ValidationError {
	return ValidationError{ErrorOf(message)}
}

func RaiseUnexpectedCommandError(command eventsourcing.Command) UnexpectedCommandError {
	return UnexpectedCommandError{ErrorOf(fmt.Sprintf("unexpected command: %#v", command))}
}

func RaiseUnexpectedEventError(event eventsourcing.Event) UnexpectedEventError {
	return UnexpectedEventError{ErrorOf(fmt.Sprintf("unexpected event: %#v", event))}
}
