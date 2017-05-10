package email

import (
	"regexp"

	"github.com/morikuni/chat/chat/domain"
	"github.com/morikuni/chat/chat/domain/model"
	"github.com/pkg/errors"
)

func Validate(email string) (model.Email, error) {
	ok, err := regexp.MatchString(`[a-zA-Z0-9\.\-\+]+@[a-zA-Z0-9\.]+`, email)
	if err != nil {
		return "", errors.Wrap(err, "failed to compile regexp")
	}
	if !ok {
		return "", errors.WithStack(domain.RaiseValidationError("invalid format"))
	}
	return model.Email(email), nil
}
