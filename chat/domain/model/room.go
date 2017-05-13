package model

import "context"

type Room interface {
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
