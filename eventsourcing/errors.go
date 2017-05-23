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
)

func RaiseUnknownEventError(typ Type) UnknowEventError {
	return UnknowEventError{ErrorOf(fmt.Sprintf("unknown event: %#v", typ))}
}

func RaiseEventVersionConflictError(event VersionedEvent) EventVersionConflictError {
	return EventVersionConflictError{ErrorOf(fmt.Sprintf("event version conflict: version=%d event=%#v", event.Version, event.Event))}
}
