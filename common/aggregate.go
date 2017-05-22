package common

type Aggregate interface {
	Changes() []VersionedEvent
	Version() uint64
	Handle(command Command) error
	Replay(events ...VersionedEvent) error
}

type aggregate struct {
	changes  []VersionedEvent
	version  uint64
	behavior Behavior
}

type Behavior interface {
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

func (a *aggregate) Changes() []VersionedEvent {
	return a.changes
}

func (a *aggregate) Version() uint64 {
	return a.version
}

func (a *aggregate) Handle(command Command) error {
	event, err := a.behavior.ReceiveCommand(command)
	if err != nil {
		return err
	}
	return a.Mutate(event)
}

func (a *aggregate) Replay(events ...VersionedEvent) error {
	for _, e := range events {
		err := a.behavior.ReceiveEvent(e.Event)
		if err != nil {
			return err
		}
		a.version = e.Version
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
		a.changes = append(a.changes, VersionedEvent{e, a.version})
	}
	return nil
}
