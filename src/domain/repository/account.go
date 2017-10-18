package repository

import (
	"context"

	"github.com/morikuni/chat/src/domain/model"
)

type Account interface {
	GenerateID(ctx context.Context) (model.UserID, error)
	Save(ctx context.Context, account *model.Account) error
	Find(ctx context.Context, id model.UserID) (*model.Account, error)
}
