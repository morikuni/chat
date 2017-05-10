package domain

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
	NoSuchEntityError struct{ DomainError }
	ValidationError   struct{ DomainError }
)

var (
	ErrNoSuchEntity = NoSuchEntityError{ErrorOf("no such entity")}
)

func RaiseValidationError(message string) ValidationError {
	return ValidationError{ErrorOf(message)}
}
