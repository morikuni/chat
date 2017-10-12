package di

import (
	"github.com/morikuni/chat/src/usecase"
	"github.com/morikuni/chat/src/usecase/impl"
)

func InjectPosting() usecase.Posting {
	return impl.NewPosting(
		InjectChatRepository(),
	)
}
