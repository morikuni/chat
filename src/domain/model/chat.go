package model

import (
	"fmt"

	"github.com/morikuni/chat/src/domain"
)

const (
	MaxMessageLength = 20
)

type ChatID int64

type ChatMessage string

func ValidateChatMessage(message string) (ChatMessage, domain.ValidationError) {
	if message == "" {
		return "", domain.RaiseValidationError("cannot be empty")
	}
	if len(message) > MaxMessageLength {
		return "", domain.RaiseValidationError(fmt.Sprintf("length must be shorter than %d", MaxMessageLength))
	}
	return ChatMessage(message), nil
}
