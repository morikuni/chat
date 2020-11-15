package usecase

import (
	"context"
	"time"

	"github.com/morikuni/failure"

	"github.com/morikuni/chat/src/domain"
)

type message struct {
	accountRepo AccountRepository
	roomRepo    RoomRepository
	messageRepo MessageRepository
	now         func() time.Time
}

var _ Message = (*message)(nil)

func NewMessage(
	accountRepo AccountRepository,
	roomRepo RoomRepository,
	messageRepo MessageRepository,
) Message {
	return &message{
		accountRepo,
		roomRepo,
		messageRepo,
		time.Now,
	}
}

type PostMessageRequest struct {
	PostedBy domain.AccountID
	RoomID   domain.RoomID
	Body     domain.MessageBody
}

type PostMessageResponse struct {
	Message *domain.Message
}

func (m *message) PostMessage(ctx context.Context, request *PostMessageRequest) (*PostMessageResponse, error) {
	account, err := m.accountRepo.Get(ctx, request.PostedBy)
	if err != nil {
		return nil, failure.Wrap(err)
	}

	room, err := m.roomRepo.Get(ctx, request.RoomID)
	if err != nil {
		return nil, failure.Wrap(err)
	}

	message := domain.NewMessage(room, account, request.Body, m.now())

	err = m.messageRepo.Save(ctx, message)
	if err != nil {
		return nil, failure.Wrap(err)
	}

	return &PostMessageResponse{message}, nil
}
