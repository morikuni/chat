package event

import (
	"time"

	"github.com/morikuni/chat/chat/domain/model"
)

type UserCreated struct {
	ID           model.UserID
	Name         model.UserName
	Email        model.Email
	PasswordHash []byte
	Salt         []byte
}

type UserProfileUpdated struct {
	Name model.UserName
}

type RoomCreated struct {
	ID          model.RoomID
	Name        model.RoomName
	Description model.RoomDescription
	OwnerID     model.UserID
	CreatedAt   time.Time
}
