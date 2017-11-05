package di

import (
	"net/http"

	"github.com/morikuni/chat/src/adapter/taskqueue"
)

func InjectTaskHandler() http.Handler {
	return taskqueue.NewTaskHandler(
		InjectLogger(),
		InjectEventHandler(),
	)
}
