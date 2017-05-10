package model

type User interface {
	ID() UserID
	Name() UserName
	Authenticate(email Email, password Password) error
	UpdateProfile(name UserName)
	JoinRoom(room Room) RoomMember
}

type UserID string

type UserName string

type UserRepository interface {
	Save(user User) error
	Find(id UserID) (User, error)
}
