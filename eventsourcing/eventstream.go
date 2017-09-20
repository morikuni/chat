package eventsourcing

type EventStream interface {
	Next() bool
	Error() error
	Event() MetaEvent
}
