package model

import "context"

type RoomMemberID string

type RoomMember interface {
	ID() RoomMemberID
	PostChat(message string) Chat
}

type RoomMemberRepository interface {
	Save(ctx context.Context, member RoomMember) error
	Find(ctx context.Context, id RoomMemberID) (RoomMember, error)
}
