package common

type Aggregate interface {
	Changes() []VersionedEvent
}

type AggregateBase struct {
	changes  []VersionedEvent
	version  uint64
	receiver EventReceiver
}

type EventReceiver interface {
	ReceiveEvent(event Event) error
}

func NewAggregateBase(receiver EventReceiver) *AggregateBase {
	return &AggregateBase{
		nil,
		0,
		receiver,
	}
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
