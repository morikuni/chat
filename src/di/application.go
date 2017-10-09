package di

import (
	"github.com/morikuni/chat/src/application"
	"github.com/morikuni/chat/src/application/read"
)

func InjectPostingService() application.PostingService {
	return application.NewPostingService(
		InjectChatRepository(),
	)
}

func InjectChatReader() read.ChatReader {
	return read.NewChatReader()
}
