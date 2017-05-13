package model

import "context"

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
	Save(ctx context.Context, user User) error
	Find(ctx context.Context, id UserID) (User, error)
}
