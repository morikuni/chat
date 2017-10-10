package reader

import (
	"context"

	"github.com/morikuni/chat/src/domain/repository"
	"github.com/morikuni/chat/src/usecase/reader/dto"
	"github.com/pkg/errors"
	"google.golang.org/appengine/datastore"
)

func NewChat() Chat {
	return chat{}
}

type Chat interface {
	Chats(ctx context.Context) ([]dto.Chat, error)
}

type chat struct{}

func (r chat) Chats(ctx context.Context) ([]dto.Chat, error) {
	var chats []dto.Chat
	_, err := datastore.NewQuery(repository.ChatKind).
		Order("-ID").
		Limit(3).
		GetAll(ctx, &chats)

	if err != nil {
		return nil, errors.Wrap(err, "failed to read chats")
	}

	return chats, nil
}
