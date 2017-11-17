package dto

import (
	"time"
)

type Chats struct {
	Chats       []Chat `json:"chats"`
	CursorToken string `json:"cursor_token"`
}

type Chat struct {
	ID       int64     `json:"id"`
	Message  string    `json:"message"`
	PostedAt time.Time `json:"posted_at"`
}
