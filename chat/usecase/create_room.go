package usecase

import (
	"context"

	"github.com/morikuni/chat/chat/domain"
	"github.com/morikuni/chat/chat/domain/model"
	"github.com/morikuni/chat/chat/domain/model/room"
	"github.com/morikuni/chat/chat/domain/model/user"
	"github.com/pkg/errors"
)

func NewCreateRoom(
	userRepository model.UserRepository,
	roomRepo model.RoomRepository,
) CreateRoom {
	return createRoom{
		userRepository,
		roomRepo,
	}
}

type CreateRoom interface {
	Create(ctx context.Context, categoryID, name, description string) (model.RoomID, error)
}

type createRoom struct {
	userRepository model.UserRepository
	roomRepo       model.RoomRepository
}

func (jr createRoom) Create(ctx context.Context, userID, name, description string) (model.RoomID, error) {
	u, err := jr.userRepository.Find(ctx, user.NewID(userID))
	if err != nil {
		switch errors.Cause(err).(type) {
		case domain.NoSuchEntityError:
			return "", errors.WithStack(ErrNoSuchUser)
		default:
			return "", errors.WithMessage(err, "failed to find category")
		}
	}
	r := u.CreateRoom(room.NewName(name), room.NewDescription(description))
	err = jr.roomRepo.Save(ctx, r)
	if err != nil {
		return "", errors.WithMessage(err, "failed to save room")
	}
	return r.ID(), nil
}
