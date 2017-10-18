package model

import (
	"regexp"

	"github.com/morikuni/chat/src/domain"
)

type Email string

func ValidateEmail(email string) (Email, domain.ValidationError) {
	if ok, _ := regexp.MatchString(`\A[a-zA-Z0-9]+@[a-zA-Z0-9\.]+\z`, email); ok {
		return Email(email), nil
	}
	return "", domain.RaiseValidationError("invalid format")
}
