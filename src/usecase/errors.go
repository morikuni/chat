package usecase

import (
	"github.com/pkg/errors"

	"github.com/morikuni/chat/src/domain"
	"github.com/morikuni/chat/src/infra"
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
)

func TranslateValidationError(err domain.ValidationError, name string) ValidationError {
	return ValidationError{
		ErrorOf(err.Error()),
		name,
	}
}
