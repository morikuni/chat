package read

import (
	"context"
	"time"

	"github.com/morikuni/chat/src/domain/repository"
	"github.com/pkg/errors"
	"google.golang.org/appengine/datastore"
)

type Chat struct {
	ID       int64
	Message  string
	PostedAt time.Time
}

func NewChatReader() ChatReader {
	return chatReader{}
}

type ChatReader interface {
	Chats(ctx context.Context) ([]Chat, error)
}

type chatReader struct{}

func (cr chatReader) Chats(ctx context.Context) ([]Chat, error) {
	var chats []Chat
	_, err := datastore.NewQuery(repository.ChatKind).
		Order("-ID").
		Limit(3).
		GetAll(ctx, &chats)

	if err != nil {
		return nil, errors.Wrap(err, "failed to read chats")
	}

	return chats, nil
}
