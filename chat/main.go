package main

import (
	"context"

	"github.com/k0kubun/pp"
	"github.com/morikuni/chat/chat/di"
	"github.com/morikuni/chat/chat/domain/model/user"
	"github.com/morikuni/chat/eventsourcing"
)

func main() {
	userRepo := di.NewUserRepository()
	eventStore := di.NewEventStore()
	ctx := context.Background()

	u := user.New("morikuni", "morikuni-taihei@kayac.com", "password")

	eventStore.Save(ctx, u.Changes())
	PrintEvents(ctx, eventStore, string(u.ID()))

	pp.Println("============================")

	u2, err := userRepo.Find(ctx, u.ID())
	if err != nil {
		panic(err)
	}
	pp.Println(u2)
	u2.UpdateProfile("taihei")

	eventStore.Save(ctx, u2.Changes())
	PrintEvents(ctx, eventStore, string(u.ID()))

	pp.Println("============================")

	u3, err := userRepo.Find(ctx, u.ID())
	if err != nil {
		panic(err)
	}
	pp.Println(u3)
}

func PrintEvents(ctx context.Context, store eventsourcing.EventStore, id string) {
	events, err := store.Find(ctx, id)
	if err != nil {
		panic(err)
	}
	pp.Println(events)
}
