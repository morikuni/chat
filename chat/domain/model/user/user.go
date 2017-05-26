package user

import (
	"github.com/morikuni/chat/chat/domain"
	"github.com/morikuni/chat/chat/domain/model"
	"github.com/morikuni/chat/chat/domain/model/roommember"
	"github.com/morikuni/chat/common"
	"github.com/morikuni/chat/eventsourcing"
	"github.com/pkg/errors"
)

func New(name model.UserName, email model.Email, password model.Password) model.User {
	u := newUser()
	u.Handle(CreateUser{name, email, password})

	return u
}

func newUser() *User {
	s := &State{}
	return &User{
		eventsourcing.NewAggregate(s),
		s,
	}
}

type User struct {
	eventsourcing.Aggregate

	state *State
}

func (u *User) ID() model.UserID {
	return u.state.id
}

func (u *User) Name() model.UserName {
	return u.state.name
}

func (u *User) Authenticate(email model.Email, password model.Password) error {
	return u.state.authInfo.Authenticate(email, password)
}

func (u *User) JoinRoom(room model.Room) model.RoomMember {
	return roommember.New(u.state.id, room.ID())
}

func (u *User) UpdateProfile(name model.UserName) {
	u.Handle(UpdateProfile{name})
}

type State struct {
	id       model.UserID
	name     model.UserName
	authInfo AuthInfo
}

func (s *State) ID() string {
	return string(s.id)
}

func (s *State) ReceiveCommand(command eventsourcing.Command) (eventsourcing.Event, error) {
	switch c := command.(type) {
	case CreateUser:
		id := common.NewUUID()
		return UserCreated{
			model.UserID(id),
			c.Name,
			newAuthInfo(c.Email, c.Password),
		}, nil
	case UpdateProfile:
		return UserProfileUpdated{
			c.Name,
		}, nil
	default:
		return nil, errors.WithStack(domain.RaiseUnexpectedCommandError(c))
	}
}

func (s *State) ReceiveEvent(event eventsourcing.Event) error {
	switch e := event.(type) {
	case UserCreated:
		s.id = e.ID
		s.name = e.Name
		s.authInfo = e.AuthInfo
	case UserProfileUpdated:
		s.name = e.Name
	default:
		return errors.WithStack(domain.RaiseUnexpectedEventError(e))
	}
	return nil
}

type CreateUser struct {
	Name     model.UserName
	Email    model.Email
	Password model.Password
}

type UserCreated struct {
	ID       model.UserID
	Name     model.UserName
	AuthInfo AuthInfo
}

type UpdateProfile struct {
	Name model.UserName
}

type UserProfileUpdated struct {
	Name model.UserName
}

func ValidateName(name string) (model.UserName, error) {
	if name == "" {
		return "", domain.RaiseValidationError("name cannot be empty")
	}
	return model.UserName(name), nil
}

func NewID(id string) model.UserID {
	return model.UserID(id)
}
