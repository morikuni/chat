package di

import (
	"github.com/morikuni/chat/src/adapter/appengine"
	"github.com/morikuni/chat/src/infra"
)

func InjectLogger() infra.Logger {
	return appengine.NewLogger()
}
