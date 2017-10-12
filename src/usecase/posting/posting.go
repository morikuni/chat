package posting

import (
	"context"

	"github.com/morikuni/chat/src/domain/model"
	"github.com/morikuni/chat/src/domain/repository"
	"github.com/morikuni/chat/src/usecase"
)

func New(chatRepository repository.Chat) usecase.Posting {
	return impl{
		chatRepository,
	}
}

type impl struct {
	chatRepository repository.Chat
}

func (ps impl) PostChat(ctx context.Context, message string) error {
	cm, verr := model.ValidateChatMessage(message)
	if verr != nil {
		return usecase.TranslateValidationError(verr, "message")
	}
	id, err := ps.chatRepository.GenerateID(ctx)
	if err != nil {
		return err
	}
	chat := model.NewChat(id, cm)
	return ps.chatRepository.Save(ctx, chat)
}
