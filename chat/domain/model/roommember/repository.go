package roommember

import (
	"github.com/morikuni/chat/chat/domain"
	"github.com/morikuni/chat/chat/domain/model"
	"sync"
)

func NewRepository() model.RoomMemberRepository {
	return &inMemoryRepo{
		make(map[model.RoomMemberID]model.RoomMember),
		sync.RWMutex{},
	}
}

type inMemoryRepo struct {
	memory map[model.RoomMemberID]model.RoomMember
	mu     sync.RWMutex
}

func (r *inMemoryRepo) Save(member model.RoomMember) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.memory[member.ID()] = member
	return nil
}

func (r *inMemoryRepo) Find(id model.RoomMemberID) (model.RoomMember, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	m, ok := r.memory[id]
	if !ok {
		return nil, domain.ErrNoSuchEntity
	}
	return m, nil
}
