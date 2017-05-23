package eventsourcing

import (
	"context"
	"fmt"
	"sync"
)

type EventCliflictError struct {
	Event VersionedEvent
}

func (e EventCliflictError) Error() string {
	return fmt.Sprintf("event version conflict: version=%d event=%#v", e.Event.Version, e.Event.Event)
}

type EventStore interface {
	Save(ctx context.Context, events []VersionedEvent) error
	Find(ctx context.Context, aggregateID string) ([]VersionedEvent, error)
}

func NewMemoryEventStore(serializer Serializer) EventStore {
	return &memoryStore{
		make(map[string][]record),
		sync.RWMutex{},
		serializer,
	}
}

type memoryStore struct {
	memory     map[string][]record
	mu         sync.RWMutex
	serializer Serializer
}

func (s *memoryStore) Save(ctx context.Context, changes []VersionedEvent) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, change := range changes {
		records := s.memory[change.Event.AggregateID()]
		for _, r := range records {
			if r.Version == change.Version {
				return RaiseEventVersionConflictError(change)
			}
		}
		j, err := s.serializer.Serialize(change.Event)
		if err != nil {
			return err
		}
		s.memory[change.Event.AggregateID()] = append(records, record{TypeOf(change.Event), change.Version, j})
	}

	return nil
}

func (s *memoryStore) Find(ctx context.Context, aggregateID string) ([]VersionedEvent, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	records := s.memory[aggregateID]
	events := make([]VersionedEvent, len(records))
	for i, r := range records {
		event, err := s.serializer.Deserialize(r.Type, r.Event)
		if err != nil {
			return nil, err
		}
		events[i] = VersionedEvent{event, r.Version}
	}
	return events, nil
}

type record struct {
	Type    Type
	Version uint64
	Event   []byte
}
