package roommember

import (
	"github.com/morikuni/chat/chat/domain/model"
	"github.com/morikuni/chat/chat/domain/model/chat"
	"github.com/morikuni/chat/common"
)

func New(userID model.UserID, roomID model.RoomID) *RoomMember {
	return &RoomMember{
		model.RoomMemberID(common.NewUUID()),
		roomID,
		userID,
	}
}

type RoomMember struct {
	id     model.RoomMemberID
	roomID model.RoomID
	userID model.UserID
}

func (rm *RoomMember) ID() model.RoomMemberID {
	return rm.id
}

func (rm *RoomMember) PostChat(message string) model.Chat {
	return chat.New(rm.roomID, message)
}
