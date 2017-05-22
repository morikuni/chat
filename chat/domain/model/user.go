package model

import (
	"context"
	"github.com/morikuni/chat/eventsourcing"
)

type User interface {
	eventsourcing.Aggregate

	ID() UserID
	Name() UserName
	Authenticate(email Email, password Password) error
	UpdateProfile(name UserName)
	JoinRoom(room Room) RoomMember
}

type UserID string

type UserName string

type UserRepository interface {
	Append(ctx context.Context, user User) error
	Save(ctx context.Context, user User) error
	Find(ctx context.Context, id UserID) (User, error)
}
