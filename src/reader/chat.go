package reader

import (
	"context"

	"github.com/morikuni/chat/src/reader/dto"
)

type Chat interface {
	Chats(ctx context.Context, cursorToken string) (dto.Chats, error)
}
