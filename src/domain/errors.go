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
	ValidationError interface {
		DomainError
		validationError()
	}
)

type validationError struct {
	DomainError
}

func (ve validationError) validationError() {}

func RaiseValidationError(message string) ValidationError {
	return validationError{ErrorOf(message)}
}
