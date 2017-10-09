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
		Parameter() string
	}
)

type validationError struct {
	DomainError
	Param string
}

func (ve validationError) Parameter() string {
	return ve.Param
}

func RaiseValidationError(param, message string) ValidationError {
	return validationError{ErrorOf(message), param}
}
