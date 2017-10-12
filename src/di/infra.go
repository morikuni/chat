package di

import (
	"github.com/morikuni/chat/src/infra"
	"github.com/morikuni/chat/src/infra/impl"
)

func InjectLogger() infra.Logger {
	return impl.NewLogger()
}
