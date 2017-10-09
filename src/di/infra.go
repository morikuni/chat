package di

import (
	"github.com/morikuni/chat/src/infra"
)

func InjectLogger() infra.Logger {
	return infra.NewLogger()
}
