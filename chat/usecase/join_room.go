package usecase

import (
	"github.com/morikuni/chat/chat/domain"
	"github.com/morikuni/chat/chat/domain/model"
	"github.com/morikuni/chat/chat/domain/model/room"
	"github.com/morikuni/chat/chat/domain/model/user"
	"github.com/pkg/errors"
)

func NewJoinRoom(
	userRepo model.UserRepository,
	roomRepo model.RoomRepository,
	roomMemberRepo model.RoomMemberRepository,
) JoinRoom {
	return joinRoom{
		userRepo,
		roomRepo,
		roomMemberRepo,
	}
}

type JoinRoom interface {
	Join(userID, roomID string) (model.RoomMemberID, error)
}

type joinRoom struct {
	userRepo       model.UserRepository
	roomRepo       model.RoomRepository
	roomMemberRepo model.RoomMemberRepository
}

func (jr joinRoom) Join(userID, roomID string) (model.RoomMemberID, error) {
	u, err := jr.userRepo.Find(user.NewID(userID))
	if err != nil {
		switch errors.Cause(err).(type) {
		case domain.NoSuchEntityError:
			return "", errors.WithStack(ErrNoSuchUser)
		default:
			return "", errors.WithMessage(err, "failed to find user")
		}
	}
	r, err := jr.roomRepo.Find(room.NewID(roomID))
	if err != nil {
		switch errors.Cause(err).(type) {
		case domain.NoSuchEntityError:
			return "", errors.WithStack(ErrNoSuchRoom)
		default:
			return "", errors.WithMessage(err, "failed to find room")
		}
	}
	rm := u.JoinRoom(r)
	err = jr.roomMemberRepo.Save(rm)
	if err != nil {
		return "", errors.WithMessage(err, "failed to save room member")
	}
	return rm.ID(), nil
}
