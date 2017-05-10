package category

import (
	"sync"

	"github.com/morikuni/chat/chat/domain"
	"github.com/morikuni/chat/chat/domain/model"
)

func NewRepository() model.CategoryRepository {
	return &inMemoryRepo{
		make(map[model.CategoryID]model.Category),
		sync.RWMutex{},
	}
}

type inMemoryRepo struct {
	memory map[model.CategoryID]model.Category
	mu     sync.RWMutex
}

func (r *inMemoryRepo) Save(category model.Category) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.memory[category.ID()] = category
	return nil
}

func (r *inMemoryRepo) Find(id model.CategoryID) (model.Category, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	category, ok := r.memory[id]
	if !ok {
		return nil, domain.ErrNoSuchEntity
	}
	return category, nil
}
