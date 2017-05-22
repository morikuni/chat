package user

import (
	"github.com/morikuni/chat/chat/domain"
	"github.com/morikuni/chat/chat/domain/model"
	"github.com/morikuni/chat/chat/domain/model/roommember"
	"github.com/morikuni/chat/common"
	"github.com/pkg/errors"
)

func New(name model.UserName, email model.Email, password model.Password) *User {
	s := &UserState{}
	u := &User{
		common.NewAggregate(s),
		s,
	}

	u.Handle(CreateUser{name, email, password})

	return u
}

type User struct {
	common.Aggregate

	state *UserState
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

type UserState struct {
	id       model.UserID
	name     model.UserName
	authInfo AuthInfo
}

func (s *UserState) ReceiveCommand(command common.Command) (common.Event, error) {
	switch c := command.(type) {
	case CreateUser:
		id := common.NewUUID()
		return UserCreated{
			common.EventOf(id),
			model.UserID(id),
			c.Name,
			newAuthInfo(c.Email, c.Password),
		}, nil
	case UpdateProfile:
		return UserProfileUpdated{
			common.EventOf(string(s.id)),
			c.Name,
		}, nil
	default:
		return nil, errors.Errorf("unexpected command: %#v", c)
	}
}

func (s *UserState) ReceiveEvent(event common.Event) error {
	switch e := event.(type) {
	case UserCreated:
		s.id = e.ID
		s.name = e.Name
		s.authInfo = e.AuthInfo
	case UserProfileUpdated:
		s.name = e.Name
	default:
		return errors.Errorf("unexpected event: %#v", e)
	}
	return nil
}

func (s *UserState) UpdateProfile(e common.EventBase, name model.UserName) common.Event {
	return UserProfileUpdated{e, name}
}

type CreateUser struct {
	Name     model.UserName
	Email    model.Email
	Password model.Password
}

type UserCreated struct {
	common.EventBase
	ID       model.UserID
	Name     model.UserName
	AuthInfo AuthInfo
}

type UpdateProfile struct {
	Name model.UserName
}

type UserProfileUpdated struct {
	common.EventBase
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
