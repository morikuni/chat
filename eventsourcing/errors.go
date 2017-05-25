package eventsourcing

import (
	"fmt"
)

type EventSourcingError interface {
	error
	eventSourcingError()
}

type eventSourcingError struct {
	message string
}

func (e eventSourcingError) Error() string {
	return e.message
}

func (e eventSourcingError) eventSourcingError() {}

func ErrorOf(message string) EventSourcingError {
	return eventSourcingError{message}
}

type (
	UnknowEventError          struct{ EventSourcingError }
	EventVersionConflictError struct{ EventSourcingError }
	NoEventsFoundError        struct{ EventSourcingError }
)

var (
	ErrNoEventsFound = NoEventsFoundError{ErrorOf("no events found")}
)

func RaiseUnknownEventError(typ Type) UnknowEventError {
	return UnknowEventError{ErrorOf(fmt.Sprintf("unknown event: %#v", typ))}
}

func RaiseEventVersionConflictError(event MetaEvent) EventVersionConflictError {
	return EventVersionConflictError{ErrorOf(fmt.Sprintf("event version conflict: meta=%#v event=%#v", event.Metadata, event.Event))}
}
