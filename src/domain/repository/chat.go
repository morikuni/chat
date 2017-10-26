package repository

import (
	"context"

	"github.com/morikuni/chat/src/domain/model"
	"github.com/morikuni/chat/src/domain/model/aggregate"
)

type Chat interface {
	GenerateID(ctx context.Context) (model.ChatID, error)
	Save(ctx context.Context, chat *aggregate.Chat) error
}
