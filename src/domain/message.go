package domain

import (
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/morikuni/chat/src/errors"
	"github.com/morikuni/failure"
)

type Message struct {
	ID       MessageID
	RoomID   RoomID
	Body     MessageBody
	PostedBy AccountID
	PostedAt time.Time
}

func NewMessage(room *Room, postedBy *Account, body MessageBody, now time.Time) *Message {
	return &Message{
		generateMessageID(),
		room.ID,
		body,
		postedBy.ID,
		now,
	}
}

type MessageID string

func generateMessageID() MessageID {
	return MessageID(uuid.New().String())
}

type MessageBody string

var messageBodyReg = regexp.MustCompile("[a-z0-9A-Z_ -]+")

func NewMessageBody(body string) (MessageBody, error) {
	if !messageBodyReg.MatchString(body) {
		return "", failure.New(errors.InvalidArgument, failure.Message("message contains invalid characters"))
	}

	if l := len(body); l == 0 || l > 140 {
		return "", failure.New(errors.InvalidArgument, failure.Message("message must be from 1 to 140 characters"))
	}

	return MessageBody(body), nil
}
