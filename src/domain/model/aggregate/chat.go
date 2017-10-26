package aggregate

import (
	"time"

	"github.com/morikuni/chat/src/domain/model"
)

type Chat struct {
	ID       model.ChatID
	Message  model.ChatMessage
	PostedAt time.Time
}

func NewChat(id model.ChatID, message model.ChatMessage, postedAt time.Time) *Chat {
	return &Chat{
		id,
		message,
		postedAt,
	}
}
