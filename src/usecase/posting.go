package usecase

import (
	"context"

	"github.com/morikuni/chat/src/domain/model"
	"github.com/morikuni/chat/src/domain/repository"
)

func NewPosting(chatRepository repository.Chat) Posting {
	return posting{
		chatRepository,
	}
}

type Posting interface {
	PostChat(ctx context.Context, message string) error
}

type posting struct {
	chatRepository repository.Chat
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
