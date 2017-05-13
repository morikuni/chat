package room

import (
	"context"
	"sync"

	"github.com/morikuni/chat/chat/domain"
	"github.com/morikuni/chat/chat/domain/model"
)

func NewRepository() model.RoomRepository {
	return &inMemoryRepo{
		make(map[model.RoomID]model.Room),
		sync.RWMutex{},
	}
}

type inMemoryRepo struct {
	memory map[model.RoomID]model.Room
	mu     sync.RWMutex
}

func (r *inMemoryRepo) Save(ctx context.Context, room model.Room) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.memory[room.ID()] = room
	return nil
}

func (r *inMemoryRepo) Find(ctx context.Context, id model.RoomID) (model.Room, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	room, ok := r.memory[id]
	if !ok {
		return nil, domain.ErrNoSuchEntity
	}
	return room, nil
}
