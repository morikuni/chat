package user

import (
	"github.com/morikuni/chat/chat/domain"
	"github.com/morikuni/chat/chat/domain/model"
	"github.com/morikuni/chat/chat/domain/model/roommember"
	"github.com/morikuni/chat/common"
)

func New(name model.UserName, email model.Email, password model.Password) *User {
	return &User{
		common.Aggregate{},

		model.UserID(common.NewUUID()),
		name,
		newAuthInfo(email, password),
	}
}

type User struct {
	common.Aggregate

	id       model.UserID
	name     model.UserName
	authInfo AuthInfo
}

func (u *User) ID() model.UserID {
	return u.id
}

func (u *User) Name() model.UserName {
	return u.name
}

func (u *User) Authenticate(email model.Email, password model.Password) error {
	return u.authInfo.Authenticate(email, password)
}

func (u *User) JoinRoom(room model.Room) model.RoomMember {
	return roommember.New(u.id, room.ID())
}

func (u *User) UpdateProfile(name model.UserName) {
	u.name = name
	u.Updated()
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
