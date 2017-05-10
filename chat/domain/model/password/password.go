package password

import (
	"github.com/morikuni/chat/chat/domain"
	"github.com/morikuni/chat/chat/domain/model"
	"github.com/pkg/errors"
)

func Validate(password string) (model.Password, error) {
	if len(password) < 6 {
		return "", errors.WithStack(domain.RaiseValidationError("password must be longer than 5 characters"))
	}
	return model.Password(password), nil
}
