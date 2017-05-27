package category

import (
	"github.com/morikuni/chat/chat/domain"
	"github.com/morikuni/chat/chat/domain/model"
	"github.com/morikuni/chat/chat/domain/model/room"
	"github.com/morikuni/chat/common"
	"github.com/morikuni/chat/eventsourcing"
	"github.com/pkg/errors"
)

func New(name model.CategoryName) model.Category {
	c := newCategory()
	err := c.Handle(CreateCategory{name})
	if err != nil {
		panic(err)
	}

	return c
}

func newCategory() *Category {
	s := &State{}
	return &Category{
		eventsourcing.NewAggregate(s),
		s,
	}
}

type Category struct {
	eventsourcing.Aggregate

	state *State
}

func (c *Category) ID() model.CategoryID {
	return c.state.id
}

func (c *Category) Name() model.CategoryName {
	return c.state.name
}

func (c *Category) AddRoom(name model.RoomName, description model.RoomDescription) model.Room {
	return room.New(c, name, description)
}

type State struct {
	id   model.CategoryID
	name model.CategoryName
}

func (s *State) ID() string {
	return string(s.id)
}

func (s *State) ReceiveCommand(command eventsourcing.Command) (eventsourcing.Event, error) {
	switch c := command.(type) {
	case CreateCategory:
		id := common.NewUUID()
		return CategoryCreated{
			model.CategoryID(id),
			c.Name,
		}, nil
	default:
		return nil, errors.WithStack(domain.RaiseUnexpectedCommandError(c))
	}
}

func (s *State) ReceiveEvent(event eventsourcing.Event) error {
	switch e := event.(type) {
	case CategoryCreated:
		s.id = e.ID
		s.name = e.Name
	default:
		return errors.WithStack(domain.RaiseUnexpectedEventError(e))
	}
	return nil
}

type CreateCategory struct {
	Name model.CategoryName
}

type CategoryCreated struct {
	ID   model.CategoryID
	Name model.CategoryName
}

func NewID(id string) model.CategoryID {
	return model.CategoryID(id)
}
