package di

import (
	"github.com/morikuni/chat/src/domain/repository"
	"github.com/morikuni/chat/src/domain/repository/chat"
)

func InjectChatRepository() repository.Chat {
	return chat.New()
}
