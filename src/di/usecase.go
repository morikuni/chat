package di

import (
	"github.com/morikuni/chat/src/usecase"
)

func InjectPosting() usecase.Posting {
	return usecase.NewPosting(
		InjectChatRepository(),
	)
}
