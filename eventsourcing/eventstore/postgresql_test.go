package eventstore

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/morikuni/chat/chat/domain/event"
	"github.com/morikuni/chat/eventsourcing"
	"github.com/stretchr/testify/assert"
)

func TestPostgresqlStore(t *testing.T) {
	assert := assert.New(t)

	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/chat?sslmode=disable")
	assert.NoError(err)

	store := postgresqlStore{db, eventsourcing.NewJSONSerializer(event.RoomCreated{})}

	events := []eventsourcing.MetaEvent{
		{
			Metadata: eventsourcing.Metadata{
				"hoge",
				time.Now(),
				1,
				eventsourcing.TypeOf(event.RoomCreated{}),
			},
			Event: event.RoomCreated{},
		},
	}

	err = store.Save(context.Background(), events)
	assert.NoError(err)

	events, err = store.Find(context.Background(), "hoge")
	assert.NoError(err)
	fmt.Printf("%#v\n", events)
}
