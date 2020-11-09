package usecase

import "context"

type Account interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
}

type Message interface {
	PostMessage(context.Context, *PostMessageRequest) (*PostMessageResponse, error)
}

type Room interface {
	CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomResponse, error)
}
