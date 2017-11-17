package di

import (
	"github.com/morikuni/chat/src/application/usecase"
)

func InjectPosting() usecase.Posting {
	return usecase.NewPosting(
		InjectChatRepository(),
		InjectClock(),
	)
}

func InjectAuthentication() usecase.Authentication {
	return usecase.NewAuthentication(
		InjectAccountRepository(),
		InjectEventPublisher(),
		InjectTransactionManager(),
	)
}

func InjectEventHandler() usecase.EventHandler {
	return usecase.NewEventHandler(
		InjectLogger(),
	)
}
