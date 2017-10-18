package usecase

import (
	"github.com/morikuni/chat/src/domain"
)

type UsecaseError interface {
	error
	usecaseError()
}

type usecaseError struct {
	message string
}

func (e usecaseError) Error() string {
	return e.message
}

func (e usecaseError) usecaseError() {}

func ErrorOf(message string) UsecaseError {
	return usecaseError{message}
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
