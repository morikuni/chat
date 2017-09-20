package usecase

import (
	"context"
	"testing"

	"github.com/morikuni/chat/eventsourcing"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	store := eventsourcing.MockEventStore{
		SaveFunc: func(ctx context.Context, events []eventsourcing.MetaEvent) error {
			return nil
		},
	}
	ru := NewRegisterUser(store)

	id, err := ru.Register(ctx, "mario", "me@email.mail", "password")
	assert.Nil(err)
	assert.NotEmpty(id)
}
