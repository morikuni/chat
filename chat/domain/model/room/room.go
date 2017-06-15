package room

import (
	"time"

	"github.com/morikuni/chat/chat/domain"
	"github.com/morikuni/chat/chat/domain/event"
	"github.com/morikuni/chat/chat/domain/model"
	"github.com/morikuni/chat/common"
	"github.com/morikuni/chat/eventsourcing"
	"github.com/pkg/errors"
)

func New(name model.RoomName, description model.RoomDescription, owner model.User) *Room {
	r := newRoom()
	err := r.Handle(CreateRoom{name, description, owner.ID()})
	if err != nil {
		panic(err)
	}

	return r
}

func newRoom() *Room {
	s := &State{}
	return &Room{
		eventsourcing.NewAggregate(s),
		s,
	}
}

type Room struct {
	eventsourcing.Aggregate

	state *State
}

func (r *Room) ID() model.RoomID {
	return r.state.id
}

func (r *Room) Name() model.RoomName {
	return r.state.name
}

func (r *Room) Description() model.RoomDescription {
	return r.state.description
}

type State struct {
	id          model.RoomID
	name        model.RoomName
	description model.RoomDescription
	ownerID     model.UserID
	createdAt   time.Time
}

func (s *State) ID() string {
	return string(s.id)
}

func (s *State) ReceiveCommand(command eventsourcing.Command) (eventsourcing.Event, error) {
	switch c := command.(type) {
	case CreateRoom:
		id := common.NewUUID()
		return event.RoomCreated{
			model.RoomID(id),
			c.Name,
			c.Description,
			c.OwnerID,
			time.Now(),
		}, nil
	default:
		return nil, errors.WithStack(domain.RaiseUnexpectedCommandError(c))
	}
}

func (s *State) ReceiveEvent(e eventsourcing.Event) error {
	switch e := e.(type) {
	case event.RoomCreated:
		s.id = e.ID
		s.name = e.Name
		s.description = e.Description
		s.ownerID = e.OwnerID
		s.createdAt = e.CreatedAt
	default:
		return errors.WithStack(domain.RaiseUnexpectedEventError(e))
	}
	return nil
}

type CreateRoom struct {
	Name        model.RoomName
	Description model.RoomDescription
	OwnerID     model.UserID
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
