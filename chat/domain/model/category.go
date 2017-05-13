package model

import "context"

type Category interface {
	ID() CategoryID
	Name() CategoryName
	AddRoom(name RoomName, description RoomDescription) Room
}

type CategoryID string

type CategoryName string

type CategoryRepository interface {
	Save(ctx context.Context, room Category) error
	Find(ctx context.Context, id CategoryID) (Category, error)
}
