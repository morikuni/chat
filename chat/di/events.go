package di

import (
	"github.com/morikuni/chat/chat/domain/event"
	"github.com/morikuni/chat/eventsourcing"
)

func NewSerializer() eventsourcing.Serializer {
	return eventsourcing.NewJSONSerializer(
		event.UserCreated{},
		event.UserProfileUpdated{},

		event.CategoryCreated{},

		event.RoomCreated{},
	)
}
