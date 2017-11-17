package datastore

import (
	"context"

	"github.com/morikuni/chat/src/domain/model"
	"github.com/morikuni/chat/src/domain/model/aggregate"
	"github.com/morikuni/chat/src/domain/repository"
	"github.com/morikuni/chat/src/reader"
	"github.com/morikuni/chat/src/reader/dto"
	"github.com/morikuni/chat/src/usecase"
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

func (chat) Save(ctx context.Context, chat *aggregate.Chat) error {
	key := datastore.NewKey(ctx, ChatKind, "", int64(chat.ID), nil)
	_, err := datastore.Put(ctx, key, chat)
	if err != nil {
		return errors.Wrap(err, "failed to save chat")
	}
	return nil
}

func (chat) Chats(ctx context.Context, cursorToken string) (dto.Chats, error) {
	q := datastore.NewQuery(ChatKind).
		Order("-PostedAt").
		Limit(3)

	var chats dto.Chats
	if cursorToken != "" {
		c, err := datastore.DecodeCursor(cursorToken)
		if err != nil {
			return chats, usecase.RaiseValidationError("cursor_token", "invalid value")
		}
		q = q.Start(c)
	}

	itr := q.Run(ctx)

	for {
		var chat dto.Chat
		_, err := itr.Next(&chat)
		if err != nil {
			if err == datastore.Done {
				break
			}
			return chats, errors.Wrap(err, "failed to read chats")
		}
		chats.Chats = append(chats.Chats, chat)
	}
	cursor, err := itr.Cursor()
	if err != nil {
		return chats, errors.Wrap(err, "failed to get cursor")
	}
	chats.CursorToken = cursor.String()

	return chats, nil
}
