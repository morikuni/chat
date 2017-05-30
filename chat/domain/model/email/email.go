package email

import (
	"regexp"

	"github.com/morikuni/chat/chat/domain"
	"github.com/morikuni/chat/chat/domain/model"
)

func Validate(email string) (model.Email, domain.ValidationError) {
	ok := regexp.MustCompile(`[a-zA-Z0-9\.\-\+]+@[a-zA-Z0-9\.]+`).MatchString(email)
	if !ok {
		return "", domain.RaiseValidationError("invalid format")
	}
	return model.Email(email), nil
}
