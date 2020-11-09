package usecase

import (
	"context"

	"github.com/morikuni/chat/src/domain"
)

type AccountRepository interface {
	Get(context.Context, domain.AccountID) (*domain.Account, error)
	Save(context.Context, *domain.Account) error
}

type MessageRepository interface {
	Save(context.Context, *domain.Message) error
}

type RoomRepository interface {
	Get(context.Context, domain.RoomID) (*domain.Room, error)
	Save(context.Context, *domain.Room) error
}
