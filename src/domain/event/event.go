package event

import (
	"github.com/morikuni/chat/src/domain/model"
)

type Event interface {
	domainEvent()
}

type AccountCreated struct {
	UserID model.UserID `json:"user_id"`
	Email  model.Email  `json:"email"`
}

func (AccountCreated) domainEvent() {}
