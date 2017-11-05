package event

import (
	"github.com/morikuni/chat/src/domain/model"
)

type Event interface {
	domainEvent()
}

type AccountCreated struct {
	UserID model.UserID
	Email  model.Email
}

func (AccountCreated) domainEvent() {}
