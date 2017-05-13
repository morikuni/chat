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
	NoSuchEntityError       struct{ DomainError }
	EntityAlreadyExistError struct{ DomainError }
	ValidationError         struct{ DomainError }
)

var (
	ErrNoSuchEntity       = NoSuchEntityError{ErrorOf("no such entity")}
	ErrEntityAlreadyExist = EntityAlreadyExistError{ErrorOf("entity already exist")}
)

func RaiseValidationError(message string) ValidationError {
	return ValidationError{ErrorOf(message)}
}
