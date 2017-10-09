package application

import (
	"fmt"

	"github.com/morikuni/chat/src/domain"
)

type ApplicationError interface {
	error
	applicationError()
}

type applicationError struct {
	message string
}

func (e applicationError) Error() string {
	return e.message
}

func (e applicationError) applicationError() {}

func ErrorOf(message string) ApplicationError {
	return applicationError{message}
}

type (
	ValidationError struct {
		ApplicationError
		Parameter string
	}
)

func TranslateValidationError(err domain.ValidationError, name string) ValidationError {
	return ValidationError{
		ErrorOf(err.Error()),
		name,
	}
}

func TranslateValidationErrorByMap(err domain.ValidationError, dict map[string]string) ValidationError {
	name, ok := dict[err.Parameter()]
	if !ok {
		panic(fmt.Sprintf("unkown parameter name %s", name))
	}
	return ValidationError{
		ErrorOf(err.Error()),
		name,
	}
}
