package model

type Category interface {
	ID() CategoryID
	Name() CategoryName
	AddRoom(name RoomName, description RoomDescription) Room
}

type CategoryID string

type CategoryName string

type CategoryRepository interface {
	Save(room Category) error
	Find(id CategoryID) (Category, error)
}
