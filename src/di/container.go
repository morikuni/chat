package di

import (
	"github.com/morikuni/chat/src/interface/http"
	"github.com/morikuni/chat/src/usecase"
)

type Container struct {
	Config *Config

	HTTPServer *http.Server

	Account usecase.Account
	Room    usecase.Room
	Message usecase.Message

	AccountRepo usecase.AccountRepository
	RoomRepo    usecase.RoomRepository
	MessageRepo usecase.MessageRepository
}

type Config struct {
	HTTPServerPortAddr string
}

func (c *Container) GetHTTPServer() *http.Server {
	if c.HTTPServer == nil {
		c.HTTPServer = http.NewServer(
			c.Config.HTTPServerPortAddr,
			c.GetAccount(),
			c.GetMessage(),
			c.GetRoom(),
		)
	}
	return c.HTTPServer
}

func (c *Container) GetAccount() usecase.Account {
	if c.Account == nil {
		c.Account = usecase.NewAccount(
			c.GetAccountRepo(),
		)
	}
	return c.Account
}

func (c *Container) GetRoom() usecase.Room {
	if c.Room == nil {
		c.Room = usecase.NewRoom(
			c.GetAccountRepo(),
			c.GetRoomRepo(),
		)
	}
	return c.Room
}

func (c *Container) GetMessage() usecase.Message {
	if c.Message == nil {
		c.Message = usecase.NewMessage(
			c.GetAccountRepo(),
			c.GetRoomRepo(),
			c.GetMessageRepo(),
		)
	}
	return c.Message
}

func (c *Container) GetAccountRepo() usecase.AccountRepository {
	if c.AccountRepo == nil {
		c.AccountRepo = nil
	}
	return c.AccountRepo
}

func (c *Container) GetRoomRepo() usecase.RoomRepository {
	if c.RoomRepo == nil {
		c.RoomRepo = nil
	}
	return c.RoomRepo
}

func (c *Container) GetMessageRepo() usecase.MessageRepository {
	if c.MessageRepo == nil {
		c.MessageRepo = nil
	}
	return c.MessageRepo
}
