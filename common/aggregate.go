package common

import "time"

type Aggregate interface {
	AggregateID() string
	Changes() []VersionedEvent
}

type AggregateBase struct {
	id       string
	changes  []VersionedEvent
	version  uint64
	receiver EventReceiver
}

type EventReceiver interface {
	ReceiveEvent(event Event) error
}

func NewAggregateBase(aggregateID string, receiver EventReceiver) *AggregateBase {
	return &AggregateBase{
		aggregateID,
		nil,
		0,
		receiver,
	}
}

func (a *AggregateBase) AggregateID() string {
	return a.id
}

func (a *AggregateBase) Changes() []VersionedEvent {
	return a.changes
}

func (a *AggregateBase) Version() uint64 {
	return a.version
}

func (a *AggregateBase) Replay(events ...VersionedEvent) error {
	for _, e := range events {
		err := a.receiver.ReceiveEvent(e.Event)
		if err != nil {
			return err
		}
		a.version = e.Version
	}
	return nil
}

func (a *AggregateBase) Mutate(events ...Event) error {
	for _, e := range events {
		err := a.receiver.ReceiveEvent(e)
		if err != nil {
			return err
		}
		a.version += 1
		a.changes = append(a.changes, VersionedEvent{e, a.version})
	}
	return nil
}

func (a *AggregateBase) EventBase() EventBase {
	return EventBase{a.id, time.Now()}
}
