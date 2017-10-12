package di

import (
	"github.com/morikuni/chat/src/usecase"
	"github.com/morikuni/chat/src/usecase/posting"
	"github.com/morikuni/chat/src/usecase/reader"
	"github.com/morikuni/chat/src/usecase/reader/chat"
)

func InjectPosting() usecase.Posting {
	return posting.New(
		InjectChatRepository(),
	)
}

func InjectChatReader() reader.Chat {
	return chat.New()
}
