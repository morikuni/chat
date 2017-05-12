package di

import (
	"github.com/morikuni/chat/chat/api"
	"github.com/morikuni/chat/chat/domain/model"
	"github.com/morikuni/chat/chat/domain/model/user"
	"github.com/morikuni/chat/chat/usecase"
)

func NewSignUp() api.SignUp {
	return api.NewSignUp(NewRegisterUser())
}

func NewRegisterUser() usecase.RegisterUser {
	return usecase.NewRegisterUser(NewUserRepository())
}

func NewUserRepository() model.UserRepository {
	return user.NewRepository()
}
