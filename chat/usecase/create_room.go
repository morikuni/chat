package usecase

import (
	"github.com/morikuni/chat/chat/domain"
	"github.com/morikuni/chat/chat/domain/model"
	"github.com/morikuni/chat/chat/domain/model/category"
	"github.com/morikuni/chat/chat/domain/model/room"
	"github.com/pkg/errors"
)

func NewCreateRoom(
	categoryRepo model.CategoryRepository,
	roomRepo model.RoomRepository,
) CreateRoom {
	return createRoom{
		categoryRepo,
		roomRepo,
	}
}

type CreateRoom interface {
	Create(categoryID, name, description string) (model.RoomID, error)
}

type createRoom struct {
	categoryRepo model.CategoryRepository
	roomRepo     model.RoomRepository
}

func (jr createRoom) Create(categoryID, name, description string) (model.RoomID, error) {
	c, err := jr.categoryRepo.Find(category.NewID(categoryID))
	if err != nil {
		switch errors.Cause(err).(type) {
		case domain.NoSuchEntityError:
			return "", errors.WithStack(ErrNoSuchCategory)
		default:
			return "", errors.WithMessage(err, "failed to find category")
		}
	}
	r := c.AddRoom(room.NewName(name), room.NewDescription(description))
	err = jr.roomRepo.Save(r)
	if err != nil {
		return "", errors.WithMessage(err, "failed to save room")
	}
	return r.ID(), nil
}
