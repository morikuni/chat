package user

import (
	"context"
	"testing"

	"github.com/morikuni/chat/chat/domain/event"
	"github.com/morikuni/chat/chat/domain/model"
	"github.com/morikuni/chat/eventsourcing"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	assert := assert.New(t)

	name := model.UserName("user")
	email := model.Email("email@email.com")
	raw := model.Password("password")
	user := New(name, email, raw)
	assert.Regexp("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}", user.ID())
	assert.Equal(name, user.Name())
	assert.Equal(eventsourcing.Version(1), user.Version())

	assert.Nil(user.Authenticate(email, raw))
	assert.NotNil(user.Authenticate(model.Email("hoge"), raw))
	assert.NotNil(user.Authenticate(email, model.Password("hoge")))

	user.UpdateProfile(model.UserName("updated"))
	assert.Equal(model.UserName("updated"), user.Name())
	assert.Equal(eventsourcing.Version(2), user.Version())

	assert.Len(user.Changes(), 2)
	assert.IsType(event.UserCreated{}, user.Changes()[0].Event)
	assert.IsType(event.UserProfileUpdated{}, user.Changes()[1].Event)

	serializer := eventsourcing.NewJSONSerializer(event.UserCreated{}, event.UserProfileUpdated{})
	store := eventsourcing.NewMemoryEventStore(serializer)
	err := store.Save(context.TODO(), user.Changes())
	assert.NoError(err)

	repo := NewRepository(store)
	user2, err := repo.Find(context.TODO(), user.ID())
	assert.NoError(err)
	assert.Equal(user.(*User).state, user2.(*User).state)
}
