package repository

import (
	"context"

	"github.com/morikuni/chat/src/domain/model"
	"github.com/pkg/errors"
	"google.golang.org/appengine/datastore"
)

const (
	ChatKind = "Chat"
)

func NewChat() Chat {
	return chat{}
}

type Chat interface {
	GenerateID(ctx context.Context) (model.ChatID, error)
	Save(ctx context.Context, chat *model.Chat) error
}

type chat struct{}

func (repo chat) GenerateID(ctx context.Context) (model.ChatID, error) {
	l, _, err := datastore.AllocateIDs(ctx, ChatKind, nil, 1)
	if err != nil {
		return 0, errors.Wrap(err, "failed to generate chat ID")
	}
	return model.ChatID(l), nil
}

func (repo chat) Save(ctx context.Context, chat *model.Chat) error {
	key := datastore.NewKey(ctx, ChatKind, "", int64(chat.ID), nil)
	_, err := datastore.Put(ctx, key, chat)
	if err != nil {
		return errors.Wrap(err, "failed to save chat")
	}
	return nil
}
