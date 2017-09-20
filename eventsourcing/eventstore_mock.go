package eventsourcing

import (
	"context"
)

var _ EventStore = MockEventStore{}

type MockEventStore struct {
	SaveFunc func(ctx context.Context, events []MetaEvent) error
	FindFunc func(ctx context.Context, aggregateID string) ([]MetaEvent, error)
}

func (mock MockEventStore) Save(ctx context.Context, events []MetaEvent) error {
	return mock.SaveFunc(ctx, events)
}

func (mock MockEventStore) Find(ctx context.Context, aggregateID string) ([]MetaEvent, error) {
	return mock.FindFunc(ctx, aggregateID)
}
