package usecase

import (
	"fmt"

	"github.com/morikuni/chat/chat/domain"
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
	ValidationError     struct{ UsecaseError }
	NoSuchUserError     struct{ UsecaseError }
	NoSuchRoomError     struct{ UsecaseError }
	NoSuchCategoryError struct{ UsecaseError }
)

var (
	ErrNoSuchUser = NoSuchUserError{ErrorOf("no such user")}
	ErrNoSuchRoom = NoSuchRoomError{ErrorOf("no such room")}
)

func RaiseValidationError(name string, err domain.ValidationError) ValidationError {
	return ValidationError{ErrorOf(fmt.Sprintf("invalid %s: %s", name, err.Error()))}
}
