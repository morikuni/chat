package category

import (
	"github.com/morikuni/chat/chat/domain/model"
	"github.com/morikuni/chat/chat/domain/model/room"
	"github.com/morikuni/chat/common"
)

func New(name model.CategoryName) model.Category {
	return &Category{
		common.Aggregate{},

		model.CategoryID(common.NewUUID()),
		name,
	}
}

type Category struct {
	common.Aggregate

	id   model.CategoryID
	name model.CategoryName
}

func (c *Category) ID() model.CategoryID {
	return c.id
}

func (c *Category) Name() model.CategoryName {
	return c.name
}

func (c *Category) AddRoom(name model.RoomName, description model.RoomDescription) model.Room {
	return room.New(c, name, description)
}

func NewID(id string) model.CategoryID {
	return model.CategoryID(id)
}
