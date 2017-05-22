package common

import (
	"time"
)

type Event interface {
	AggregateID() string
	OccuredAt() time.Time
}

type VersionedEvent struct {
	Event   Event
	Version uint64
}

type EventBase struct {
	ID string
	At time.Time
}

func (e EventBase) AggregateID() string {
	return e.ID
}

func (e EventBase) OccuredAt() time.Time {
	return e.At
}
