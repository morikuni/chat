package usecase

import (
	"context"

	"github.com/morikuni/chat/src/domain/model"
)

func NewPosting(chatRepository model.ChatRepository) Posting {
	return posting{
		chatRepository,
	}
}

type Posting interface {
	PostChat(ctx context.Context, message string) error
}

type posting struct {
	chatRepository model.ChatRepository
}

func (ps posting) PostChat(ctx context.Context, message string) error {
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
