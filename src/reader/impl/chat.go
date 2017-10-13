package impl

import (
	"context"

	repo "github.com/morikuni/chat/src/domain/repository/impl"
	"github.com/morikuni/chat/src/reader"
	"github.com/morikuni/chat/src/reader/dto"
	"github.com/pkg/errors"
	"google.golang.org/appengine/datastore"
)

func NewChat() reader.Chat {
	return impl{}
}

type impl struct{}

func (r impl) Chats(ctx context.Context) ([]dto.Chat, error) {
	var chats []dto.Chat
	_, err := datastore.NewQuery(repo.ChatKind).
		Order("-PostedAt").
		Limit(3).
		GetAll(ctx, &chats)

	if err != nil {
		return nil, errors.Wrap(err, "failed to read chats")
	}

	return chats, nil
}
