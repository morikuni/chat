package di

import (
	"github.com/morikuni/chat/src/domain/repository"
	"github.com/morikuni/chat/src/domain/repository/impl"
)

func InjectChatRepository() repository.Chat {
	return impl.NewChat()
}
