package eventsourcing

import (
	"time"
)

type Command interface{}

type Event interface{}

type Version uint64

type MetaEvent struct {
	Metadata Metadata
	Event    Event
}

type Metadata struct {
	AggregateID string
	OccuredAt   time.Time
	Version     Version
	Type        Type
}
