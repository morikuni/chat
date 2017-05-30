package user

import (
	"github.com/morikuni/chat/chat/domain"
	"github.com/morikuni/chat/chat/domain/event"
	"github.com/morikuni/chat/chat/domain/model"
	"github.com/morikuni/chat/chat/domain/model/roommember"
	"github.com/morikuni/chat/common"
	"github.com/morikuni/chat/eventsourcing"
	"github.com/pkg/errors"
)

func New(name model.UserName, email model.Email, password model.Password) model.User {
	u := newUser()
	err := u.Handle(CreateUser{name, email, password})
	if err != nil {
		panic(err)
	}

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
	err := u.Handle(UpdateProfile{name})
	if err != nil {
		panic(err)
	}
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
		authInfo := newAuthInfo(c.Email, c.Password)
		return event.UserCreated{
			model.UserID(id),
			c.Name,
			authInfo.Email,
			authInfo.Password.Hash,
			authInfo.Password.Salt,
		}, nil
	case UpdateProfile:
		return event.UserProfileUpdated{
			c.Name,
		}, nil
	default:
		return nil, errors.WithStack(domain.RaiseUnexpectedCommandError(c))
	}
}

func (s *State) ReceiveEvent(e eventsourcing.Event) error {
	switch e := e.(type) {
	case event.UserCreated:
		s.id = e.ID
		s.name = e.Name
		s.authInfo.Email = e.Email
		s.authInfo.Password.Hash = e.PasswordHash
		s.authInfo.Password.Salt = e.Salt
	case event.UserProfileUpdated:
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

type UpdateProfile struct {
	Name model.UserName
}

func ValidateName(name string) (model.UserName, domain.ValidationError) {
	if name == "" {
		return "", domain.RaiseValidationError("name cannot be empty")
	}
	return model.UserName(name), nil
}

func NewID(id string) model.UserID {
	return model.UserID(id)
}
