package di

import (
	"github.com/morikuni/chat/src/adapter/datastore"
	"github.com/morikuni/chat/src/adapter/taskqueue"
	"github.com/morikuni/chat/src/domain/event"
	"github.com/morikuni/chat/src/domain/repository"
)

func InjectChatRepository() repository.Chat {
	return datastore.NewChatRepository()
}

func InjectEventPublisher() event.Publisher {
	return taskqueue.NewEventPublisher()
}
