package dto

import (
	"time"
)

type Chat struct {
	ID       int64
	Message  string
	PostedAt time.Time
}
