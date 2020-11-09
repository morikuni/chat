package domain

import (
	"regexp"

	"github.com/google/uuid"
	"github.com/morikuni/chat/src/errors"
	"github.com/morikuni/failure"
)

type Room struct {
	ID        RoomID
	Name      RoomName
	CreatedBy AccountID
}

func NewRoom(createdBy *Account, name RoomName) *Room {
	return &Room{
		generateRoomID(),
		name,
		createdBy.ID,
	}
}

type RoomID string

func NewRoomID(id string) (RoomID, error) {
	if id == "" {
		return "", failure.New(errors.InvalidArgument,
			failure.Message("room id must not be empty"),
		)
	}

	return RoomID(id), nil
}

func generateRoomID() RoomID {
	return RoomID(uuid.New().String())
}

type RoomName string

var roomNameReg = regexp.MustCompile("[a-z0-9A-Z_-]+")

func NewRoomName(name string) (RoomName, error) {
	if !roomNameReg.MatchString(name) {
		return "", failure.New(errors.InvalidArgument, failure.Message("room name contains invalid characters"))
	}

	if l := len(name); l == 0 || l > 20 {
		return "", failure.New(errors.InvalidArgument, failure.Message("room name must be from 1 to 140 characters"))
	}

	return RoomName(name), nil
}
