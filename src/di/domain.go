package di

import (
	"github.com/morikuni/chat/src/domain/repository"
)

func InjectChatRepository() repository.Chat {
	return repository.NewChatRepository()
}
