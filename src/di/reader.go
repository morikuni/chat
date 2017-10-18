package di

import (
	"github.com/morikuni/chat/src/adapter/datastore"
	"github.com/morikuni/chat/src/reader"
)

func InjectChatReader() reader.Chat {
	return datastore.NewChatReader()
}
