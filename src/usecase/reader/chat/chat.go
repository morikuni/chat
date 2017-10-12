package chat

import (
	"context"

	"github.com/morikuni/chat/src/domain/repository/chat"
	"github.com/morikuni/chat/src/usecase/reader"
	"github.com/morikuni/chat/src/usecase/reader/dto"
	"github.com/pkg/errors"
	"google.golang.org/appengine/datastore"
)

func New() reader.Chat {
	return impl{}
}

type impl struct{}

func (r impl) Chats(ctx context.Context) ([]dto.Chat, error) {
	var chats []dto.Chat
	_, err := datastore.NewQuery(chat.Kind).
		Order("-PostedAt").
		Limit(3).
		GetAll(ctx, &chats)

	if err != nil {
		return nil, errors.Wrap(err, "failed to read chats")
	}

	return chats, nil
}