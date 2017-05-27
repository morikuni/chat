package di

import (
	"github.com/morikuni/chat/chat/api"
	"github.com/morikuni/chat/chat/domain/model"
	"github.com/morikuni/chat/chat/domain/model/user"
	"github.com/morikuni/chat/chat/usecase"
	"github.com/morikuni/chat/eventsourcing"
)

func NewSignUp() api.SignUp {
	return api.NewSignUp(NewRegisterUser())
}

func NewRegisterUser() usecase.RegisterUser {
	return usecase.NewRegisterUser(NewEventStore(), NewUserRepository())
}

func NewUserRepository() model.UserRepository {
	return user.NewRepository(NewEventStore())
}

var eventStore = eventsourcing.NewMemoryEventStore(NewSerializer())

func NewEventStore() eventsourcing.EventStore {
	return eventStore
}
