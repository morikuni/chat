package di

import (
	"github.com/morikuni/chat/src/usecase"
	"github.com/morikuni/chat/src/usecase/reader"
)

func InjectPosting() usecase.Posting {
	return usecase.NewPosting(
		InjectChatRepository(),
	)
}

func InjectChatReader() reader.Chat {
	return reader.NewChat()
}
