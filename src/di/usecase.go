package di

import (
	"github.com/morikuni/chat/src/usecase"
	"github.com/morikuni/chat/src/usecase/read"
)

func InjectPosting() usecase.Posting {
	return usecase.NewPosting(
		InjectChatRepository(),
	)
}

func InjectChatReader() read.ChatReader {
	return read.NewChatReader()
}
