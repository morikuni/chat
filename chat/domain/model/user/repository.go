package user

import (
	"context"
	"sync"

	"github.com/morikuni/chat/chat/domain"
	"github.com/morikuni/chat/chat/domain/model"
)

func NewRepository() model.UserRepository {
	return &inMemoryRepo{
		make(map[model.UserID]model.User),
		sync.RWMutex{},
	}
}

type inMemoryRepo struct {
	memory map[model.UserID]model.User
	mu     sync.RWMutex
}

func (r *inMemoryRepo) Save(ctx context.Context, user model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.memory[user.ID()] = user
	return nil
}

func (r *inMemoryRepo) Find(ctx context.Context, id model.UserID) (model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	u, ok := r.memory[id]
	if !ok {
		return nil, domain.ErrNoSuchEntity
	}
	return u, nil
}
