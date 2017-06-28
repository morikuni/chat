package eventstore

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
	"github.com/morikuni/chat/eventsourcing"
	"github.com/pkg/errors"
)

func NewPostgresqlStore(db *sql.DB, serializer eventsourcing.Serializer) eventsourcing.EventStore {
	return &postgresqlStore{db, serializer}
}

type postgresqlStore struct {
	db         *sql.DB
	serializer eventsourcing.Serializer
}

func (s *postgresqlStore) Save(ctx context.Context, changes []eventsourcing.MetaEvent) error {
	tx, err := s.db.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}
	defer tx.Rollback()

	stmt, err := s.db.PrepareContext(ctx, `
		INSERT INTO event (aggregate_id, version, package, type, occured_at, data) values ($1, $2, $3, $4, $5, $6)
	`)
	if err != nil {
		return errors.Wrap(err, "failed to prepare statement")
	}
	defer stmt.Close()

	for _, event := range changes {
		data, err := s.serializer.Serialize(event.Event)
		if err != nil {
			return errors.WithMessage(err, "failed to serialize event")
		}

		meta := event.Metadata
		_, err = stmt.Exec(meta.AggregateID, meta.Version, meta.Type.Package, meta.Type.Name, meta.OccuredAt, data)
		if err != nil {
			if perr, ok := err.(*pq.Error); ok && perr.Code == "23505" {
				return errors.WithStack(eventsourcing.RaiseEventVersionConflictError(event))
			}
			return errors.Wrap(err, "failed to execute statement")
		}
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}

	return nil
}

func (s *postgresqlStore) Find(ctx context.Context, aggregateID string) ([]eventsourcing.MetaEvent, error) {
	rows, err := s.db.QueryContext(ctx, `
		SELECT version, package, type, occured_at, data FROM event WHERE aggregate_id = $1 ORDER BY version ASC
	`, aggregateID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query")
	}

	var events []eventsourcing.MetaEvent
	for rows.Next() {
		var version eventsourcing.Version
		var pkg, typ string
		var occuredAt time.Time
		var data []byte

		err := rows.Scan(&version, &pkg, &typ, &occuredAt, &data)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan rows")
		}

		t := eventsourcing.Type{
			pkg,
			typ,
		}
		event, err := s.serializer.Deserialize(t, data)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to deserialize data")
		}

		events = append(events, eventsourcing.MetaEvent{
			eventsourcing.Metadata{
				aggregateID,
				occuredAt,
				version,
				t,
			},
			event,
		})
	}

	if len(events) == 0 {
		return nil, errors.WithStack(eventsourcing.ErrNoEventsFound)
	}

	return events, nil
}
