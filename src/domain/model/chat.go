package model

import (
	"context"
	"fmt"
	"time"

	"github.com/morikuni/chat/src/domain"
	"github.com/pkg/errors"
	"google.golang.org/appengine/datastore"
)

const (
	MaxMessageLength = 20
	ChatKind         = "Chat"
)

type Chat struct {
	ID       ChatID
	Message  ChatMessage
	PostedAt time.Time
}

func NewChat(id ChatID, message ChatMessage) *Chat {
	return &Chat{
		id,
		message,
		time.Now(),
	}
}

type ChatID int64

type ChatMessage string

func ValidateChatMessage(message string) (ChatMessage, domain.ValidationError) {
	if message == "" {
		return "", domain.RaiseValidationError("message", "cannot be empty")
	}
	if len(message) > MaxMessageLength {
		return "", domain.RaiseValidationError("message", fmt.Sprintf("length must be shorter than %d", MaxMessageLength))
	}
	return ChatMessage(message), nil
}

func NewChatRepository() ChatRepository {
	return chatRepository{}
}

type ChatRepository interface {
	GenerateID(ctx context.Context) (ChatID, error)
	Save(ctx context.Context, chat *Chat) error
}

type chatRepository struct{}

func (repo chatRepository) GenerateID(ctx context.Context) (ChatID, error) {
	l, _, err := datastore.AllocateIDs(ctx, ChatKind, nil, 1)
	if err != nil {
		return 0, errors.Wrap(err, "failed to generate chat ID")
	}
	return ChatID(l), nil
}

func (repo chatRepository) Save(ctx context.Context, chat *Chat) error {
	key := datastore.NewKey(ctx, ChatKind, "", int64(chat.ID), nil)
	_, err := datastore.Put(ctx, key, chat)
	if err != nil {
		return errors.Wrap(err, "failed to save chat")
	}
	return nil
}
