package usecase

import (
	"context"

	"github.com/morikuni/chat/src/domain"
	"github.com/morikuni/failure"
)

type room struct {
	accountRepo AccountRepository
	roomRepo    RoomRepository
}

var _ Room = (*room)(nil)

func NewRoom(
	accountRepo AccountRepository,
	roomRepo RoomRepository,
) Room {
	return &room{
		accountRepo,
		roomRepo,
	}
}

type CreateRoomRequest struct {
	OwnerID domain.AccountID
	Name    domain.RoomName
}

type CreateRoomResponse struct {
	Room *domain.Room
}

func (r *room) CreateRoom(ctx context.Context, request *CreateRoomRequest) (*CreateRoomResponse, error) {
	account, err := r.accountRepo.Get(ctx, request.OwnerID)
	if err != nil {
		return nil, failure.Wrap(err)
	}

	room := domain.NewRoom(account, request.Name)

	err = r.roomRepo.Save(ctx, room)
	if err != nil {
		return nil, failure.Wrap(err)
	}

	return &CreateRoomResponse{room}, nil
}
