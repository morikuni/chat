package api

import (
	"net/http"

	"github.com/morikuni/chat/chat/usecase"
	"github.com/pkg/errors"
)

type SignUp struct {
	registerUser usecase.RegisterUser
}

func (api SignUp) Path() string {
	return "/signup"
}

func (api SignUp) TryServeHTTP(w http.ResponseWriter, r *http.Request) error {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	userID, err := api.registerUser.Register(name, email, password)
	if err != nil {
		return errors.WithMessage(err, "failed to register user")
	}

	return JSON(w, map[string]interface{}{
		"id": userID,
	})
}
