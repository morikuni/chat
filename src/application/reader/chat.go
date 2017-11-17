package reader

import (
	"context"

	"github.com/morikuni/chat/src/application/dto"
)

type Chat interface {
	Chats(ctx context.Context, cursorToken string) (dto.Chats, error)
}
