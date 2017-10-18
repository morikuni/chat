package usecase

import (
	"context"

	"github.com/morikuni/chat/src/domain/model"
	"github.com/morikuni/chat/src/domain/repository"
)

type Posting interface {
	PostChat(ctx context.Context, message string) error
}

func NewPosting(chatRepository repository.Chat) Posting {
	return posting{
		chatRepository,
	}
}

type posting struct {
	chatRepository repository.Chat
}

func (p posting) PostChat(ctx context.Context, message string) error {
	cm, verr := model.ValidateChatMessage(message)
	if verr != nil {
		return TranslateValidationError(verr, "message")
	}
	id, err := p.chatRepository.GenerateID(ctx)
	if err != nil {
		return err
	}
	chat := model.NewChat(id, cm)
	return p.chatRepository.Save(ctx, chat)
}
