package application

import (
	"context"

	"github.com/morikuni/chat/src/domain/model"
)

func NewPostingService(chatRepository model.ChatRepository) PostingService {
	return postingService{
		chatRepository,
	}
}

type PostingService interface {
	PostChat(ctx context.Context, message string) error
}

type postingService struct {
	chatRepository model.ChatRepository
}

func (ps postingService) PostChat(ctx context.Context, message string) error {
	cm, verr := model.ValidateChatMessage(message)
	if verr != nil {
		return TranslateValidationError(verr, "message")
	}
	id, err := ps.chatRepository.GenerateID(ctx)
	if err != nil {
		return err
	}
	chat := model.NewChat(id, cm)
	return ps.chatRepository.Save(ctx, chat)
}
