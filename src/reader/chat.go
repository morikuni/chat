package reader

import (
	"context"

	"github.com/morikuni/chat/src/reader/dto"
)

type Chat interface {
	Chats(ctx context.Context) ([]dto.Chat, error)
}
