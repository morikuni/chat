package model

import (
	"time"
)

type Chat interface {
	Message() string
	PostedAt() time.Time
}
