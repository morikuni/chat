package eventsourcing

type EventSource interface {
	Stream(aggregateID string) EventStream
	StreamSince(aggregateID string, since Version) EventStream
}
