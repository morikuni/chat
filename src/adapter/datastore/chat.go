package datastore

import (
	"context"

	"github.com/morikuni/chat/src/domain/model"
	"github.com/morikuni/chat/src/domain/repository"
	"github.com/morikuni/chat/src/reader"
	"github.com/morikuni/chat/src/reader/dto"
	"github.com/pkg/errors"
	"google.golang.org/appengine/datastore"
)

const (
	ChatKind = "Chat"
)

func NewChatRepository() repository.Chat {
	return chat{}
}

func NewChatReader() reader.Chat {
	return chat{}
}

type chat struct{}

func (chat) GenerateID(ctx context.Context) (model.ChatID, error) {
	l, _, err := datastore.AllocateIDs(ctx, ChatKind, nil, 1)
	if err != nil {
		return 0, errors.Wrap(err, "failed to generate chat ID")
	}
	return model.ChatID(l), nil
}

func (chat) Save(ctx context.Context, chat *model.Chat) error {
	key := datastore.NewKey(ctx, ChatKind, "", int64(chat.ID), nil)
	_, err := datastore.Put(ctx, key, chat)
	if err != nil {
		return errors.Wrap(err, "failed to save chat")
	}
	return nil
}

func (chat) Chats(ctx context.Context) ([]dto.Chat, error) {
	var chats []dto.Chat
	_, err := datastore.NewQuery(ChatKind).
		Order("-PostedAt").
		Limit(3).
		GetAll(ctx, &chats)

	if err != nil {
		return nil, errors.Wrap(err, "failed to read chats")
	}

	return chats, nil
}
