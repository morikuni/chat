package di

import (
	"database/sql"
	"github.com/morikuni/chat/chat/api"
	"github.com/morikuni/chat/chat/domain/model"
	"github.com/morikuni/chat/chat/domain/model/user"
	"github.com/morikuni/chat/chat/usecase"
	"github.com/morikuni/chat/eventsourcing"

	_ "github.com/lib/pq"
	"github.com/morikuni/chat/eventsourcing/eventstore"
)

func NewSignUp() api.SignUp {
	return api.NewSignUp(NewRegisterUser())
}

func NewRegisterUser() usecase.RegisterUser {
	return usecase.NewRegisterUser(NewEventStore())
}

func NewUserRepository() model.UserRepository {
	return user.NewRepository(NewEventStore())
}

func NewEventStore() eventsourcing.EventStore {
	db, err := sql.Open("postgres", "postgresql://root@localhost:26257?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return eventstore.NewPostgresqlStore(db, NewSerializer())
}
