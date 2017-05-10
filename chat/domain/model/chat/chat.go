package chat

import (
	"github.com/morikuni/chat/chat/domain/model"
	"time"
)

func New(roomID model.RoomID, message string) Chat {
	return Chat{
		roomID,
		message,
		time.Now(),
	}
}

type Chat struct {
	roomID   model.RoomID
	message  string
	postedAt time.Time
}

func (c Chat) Message() string {
	return c.message
}

func (c Chat) PostedAt() time.Time {
	return c.postedAt
}
