package api

import (
	"github.com/morikuni/chat/src/application"
)

type Error struct {
	Error string      `json:"error"`
	Info  interface{} `json:"info,omitempty"`
}

var (
	InternalServerError = Error{
		"internal server error",
		nil,
	}
)

func ValidationError(err application.ValidationError) Error {
	return Error{
		"validation error",
		struct {
			Parameter string
			Reason    string
		}{
			err.Parameter,
			err.Error(),
		},
	}
}
