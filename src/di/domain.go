package di

import (
	"github.com/morikuni/chat/src/domain/model"
)

func InjectChatRepository() model.ChatRepository {
	return model.NewChatRepository()
}
