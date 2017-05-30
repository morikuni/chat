package password

import (
	"github.com/morikuni/chat/chat/domain"
	"github.com/morikuni/chat/chat/domain/model"
)

func Validate(password string) (model.Password, domain.ValidationError) {
	if len(password) < 6 {
		return "", domain.RaiseValidationError("password must be longer than 5 characters")
	}
	return model.Password(password), nil
}
