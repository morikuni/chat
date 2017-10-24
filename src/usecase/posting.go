package usecase

import (
	"context"

	"github.com/morikuni/chat/src/domain/model"
	"github.com/morikuni/chat/src/domain/repository"
	"github.com/morikuni/chat/src/infra"
)

type Posting interface {
	PostChat(ctx context.Context, message string) error
}

func NewPosting(
	chatRepository repository.Chat,
	clock infra.Clock,
) Posting {
	return posting{
		chatRepository,
		clock,
	}
}

type posting struct {
	chatRepository repository.Chat
	clock          infra.Clock
}

func (p posting) PostChat(ctx context.Context, message string) error {
	cm, verr := model.ValidateChatMessage(message)
	if verr != nil {
		return verr
	}
	id, err := p.chatRepository.GenerateID(ctx)
	if err != nil {
		return err
	}
	chat := model.NewChat(id, cm, p.clock.Now())
	return p.chatRepository.Save(ctx, chat)
}
