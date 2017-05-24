package eventsourcing

import (
	"time"
)

type Aggregate interface {
	Changes() []MetaEvent
	Version() Version
	Handle(command Command) error
	Replay(events []MetaEvent) error
}

type aggregate struct {
	changes  []MetaEvent
	version  Version
	behavior Behavior
}

type Behavior interface {
	ID() string
	ReceiveCommand(command Command) (Event, error)
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

func (a *aggregate) Handle(command Command) error {
	event, err := a.behavior.ReceiveCommand(command)
	if err != nil {
		return err
	}
	return a.Mutate(event)
}

func (a *aggregate) Replay(events []MetaEvent) error {
	for _, e := range events {
		err := a.behavior.ReceiveEvent(e.Data)
		if err != nil {
			return err
		}
		a.version = e.Metadata.Version
	}
	return nil
}

func (a *aggregate) Mutate(events ...Event) error {
	for _, e := range events {
		err := a.behavior.ReceiveEvent(e)
		if err != nil {
			return err
		}
		a.version += 1
		a.changes = append(a.changes, MetaEvent{
			Metadata{
				a.behavior.ID(),
				time.Now(),
				a.version,
				TypeOf(e),
			},
			e,
		})
	}
	return nil
}
