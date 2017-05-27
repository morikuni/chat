package model

import (
	"context"

	"github.com/morikuni/chat/eventsourcing"
)

type Room interface {
	eventsourcing.Aggregate

	ID() RoomID
	Name() RoomName
	Description() RoomDescription
}

type RoomID string

type RoomName string

type RoomDescription string

type RoomRepository interface {
	Save(ctx context.Context, room Room) error
	Find(ctx context.Context, id RoomID) (Room, error)
}
