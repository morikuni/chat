package application

import (
	"fmt"

	"github.com/morikuni/chat/src/domain"
	"github.com/morikuni/chat/src/domain/event"
	"github.com/morikuni/chat/src/infra"
	"github.com/pkg/errors"
)

type UsecaseError interface {
	infra.StackTraceError
	usecaseError()
}

type usecaseError struct {
	infra.StackTraceError
}

func (e usecaseError) usecaseError() {}

func ErrorOf(message string) UsecaseError {
	return usecaseError{errors.New(message).(infra.StackTraceError)}
}

type (
	ValidationError struct {
		UsecaseError
		Parameter string
	}
	UnknownEventError struct {
		UsecaseError
		Event event.Event
	}
	InvalidCredentialError struct {
		UsecaseError
	}
)

func RaiseValidationError(name, reason string) error {
	return ValidationError{
		ErrorOf(reason),
		name,
	}
}

func TranslateValidationError(err domain.ValidationError, name string) error {
	return ValidationError{
		ErrorOf(err.Error()),
		name,
	}
}

func RaiseUnknownEventError(e event.Event) error {
	return UnknownEventError{
		ErrorOf(fmt.Sprintf("unknown event: %#v", e)),
		e,
	}
}

func RaiseInvalidCredentialError() error {
	return InvalidCredentialError{
		ErrorOf("invalid credential"),
	}
}
