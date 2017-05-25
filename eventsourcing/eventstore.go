package eventsourcing

import (
	"context"
	"sync"

	"github.com/pkg/errors"
)

type EventStore interface {
	Save(ctx context.Context, events []MetaEvent) error
	Find(ctx context.Context, aggregateID string) ([]MetaEvent, error)
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

func (s *memoryStore) Save(ctx context.Context, changes []MetaEvent) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, change := range changes {
		records := s.memory[change.Metadata.AggregateID]
		for _, r := range records {
			if r.Metadata.Version == change.Metadata.Version {
				return errors.WithStack(RaiseEventVersionConflictError(change))
			}
		}
		data, err := s.serializer.Serialize(change.Event)
		if err != nil {
			return errors.WithMessage(err, "failed to serialize event")
		}
		s.memory[change.Metadata.AggregateID] = append(records, record{change.Metadata, data})
	}

	return nil
}

func (s *memoryStore) Find(ctx context.Context, aggregateID string) ([]MetaEvent, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	records := s.memory[aggregateID]
	if len(records) == 0 {
		return nil, ErrNoEventsFound
	}

	events := make([]MetaEvent, len(records))
	for i, r := range records {
		event, err := s.serializer.Deserialize(r.Metadata.Type, r.Data)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to deserialize data")
		}
		events[i] = MetaEvent{r.Metadata, event}
	}
	return events, nil
}

type record struct {
	Metadata Metadata
	Data     []byte
}
