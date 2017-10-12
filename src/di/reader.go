package di

import (
	"github.com/morikuni/chat/src/usecase/reader"
	"github.com/morikuni/chat/src/usecase/reader/impl"
)

func InjectChatReader() reader.Chat {
	return impl.NewChat()
}
