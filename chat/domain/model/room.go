package model

type Room interface {
	ID() RoomID
	Name() RoomName
	Description() RoomDescription
}

type RoomID string

type RoomName string

type RoomDescription string

type RoomRepository interface {
	Save(room Room) error
	Find(id RoomID) (Room, error)
}
