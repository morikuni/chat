package room

import (
	"time"

	"github.com/morikuni/chat/chat/domain/model"
	"github.com/morikuni/chat/common"
)

func New(category model.Category, name model.RoomName, description model.RoomDescription) *Room {
	return &Room{
		common.Aggregate{},

		model.RoomID(common.NewUUID()),
		name,
		description,
		category.ID(),
		time.Now(),
	}
}

type Room struct {
	common.Aggregate

	id          model.RoomID
	name        model.RoomName
	description model.RoomDescription
	categoryID  model.CategoryID
	createdAt   time.Time
}

func (r *Room) ID() model.RoomID {
	return r.id
}

func (r *Room) Name() model.RoomName {
	return r.name
}

func (r *Room) Description() model.RoomDescription {
	return r.description
}

func NewID(id string) model.RoomID {
	return model.RoomID(id)
}

func NewName(name string) model.RoomName {
	return model.RoomName(name)
}

func NewDescription(description string) model.RoomDescription {
	return model.RoomDescription(description)
}
