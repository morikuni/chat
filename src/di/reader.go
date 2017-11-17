package di

import (
	"github.com/morikuni/chat/src/adapter/datastore"
	"github.com/morikuni/chat/src/application/reader"
)

func InjectChatReader() reader.Chat {
	return datastore.NewChatReader()
}
