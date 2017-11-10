package di

import (
	"github.com/morikuni/chat/src/adapter/appengine"
	"github.com/morikuni/chat/src/adapter/datastore"
	"github.com/morikuni/chat/src/infra"
)

func InjectLogger() infra.Logger {
	return appengine.NewLogger()
}

func InjectClock() infra.Clock {
	return infra.NewClock()
}

func InjectTransactionManager() infra.TransactionManager {
	return datastore.NewTransactionManager()
}
