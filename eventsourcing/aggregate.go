package eventsourcing

import (
	"time"
)

type Aggregate interface {
	Changes() []MetaEvent
	Version() Version
	Apply(event Event) error
	Replay(events []MetaEvent) error
}

type aggregate struct {
	changes  []MetaEvent
	version  Version
	behavior Behavior
}

type Behavior interface {
	ID() string
	ReceiveEvent(event Event) error
}

func NewAggregate(behavior Behavior) Aggregate {
	return &aggregate{
		nil,
		0,
		behavior,
	}
}

func (a *aggregate) Changes() []MetaEvent {
	return a.changes
}

func (a *aggregate) Version() Version {
	return a.version
}

func (a *aggregate) Apply(event Event) error {
	err := a.behavior.ReceiveEvent(event)
	if err != nil {
		return err
	}
	a.version += 1
	a.changes = append(a.changes, MetaEvent{
		Metadata{
			a.behavior.ID(),
			time.Now(),
			a.version,
			TypeOf(event),
		},
		event,
	})
	return nil
}

func (a *aggregate) Replay(events []MetaEvent) error {
	for _, e := range events {
		err := a.behavior.ReceiveEvent(e.Event)
		if err != nil {
			return err
		}
		a.version = e.Metadata.Version
	}
	return nil
}
