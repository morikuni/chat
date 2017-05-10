package model

type RoomMemberID string

type RoomMember interface {
	ID() RoomMemberID
	PostChat(message string) Chat
}

type RoomMemberRepository interface {
	Save(member RoomMember) error
	Find(id RoomMemberID) (RoomMember, error)
}
