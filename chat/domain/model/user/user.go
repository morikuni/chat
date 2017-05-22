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
	id := common.NewUUID()
	u := &User{
		common.NewAggregateBase(id, s),
		s,
	}

	event := UserCreated{
		u.EventBase(),
		model.UserID(id),
		name,
		newAuthInfo(email, password),
	}

	u.Mutate(event)

	return u
}

type User struct {
	*common.AggregateBase

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
	u.Mutate(u.state.UpdateProfile(u.EventBase(), name))
}

type UserState struct {
	id       model.UserID
	name     model.UserName
	authInfo AuthInfo
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

type UserCreated struct {
	common.EventBase
	ID       model.UserID
	Name     model.UserName
	AuthInfo AuthInfo
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
