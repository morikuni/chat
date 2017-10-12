package repository

import (
	"context"

	"github.com/morikuni/chat/src/domain/model"
)

type Chat interface {
	GenerateID(ctx context.Context) (model.ChatID, error)
	Save(ctx context.Context, chat *model.Chat) error
}
