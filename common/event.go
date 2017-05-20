package common

import (
	"time"
)

type Event interface {
	AggregateID() string
	OccuredAt() time.Time
}

type VersionedEvent struct {
	Event
	Version uint64
}

type EventBase struct {
	aggregateID string
	occuredAt   time.Time
}

func (e EventBase) AggregateID() string {
	return e.aggregateID
}

func (e EventBase) OccuredAt() time.Time {
	return e.occuredAt
}

func EventOf(id string) EventBase {
	return EventBase{id, time.Now()}
}
