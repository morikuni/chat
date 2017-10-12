package di

import (
	"github.com/morikuni/chat/src/reader"
	"github.com/morikuni/chat/src/reader/impl"
)

func InjectChatReader() reader.Chat {
	return impl.NewChat()
}
