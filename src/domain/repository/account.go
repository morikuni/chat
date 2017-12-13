package repository

import (
	"context"

	"github.com/morikuni/chat/src/domain/model"
	"github.com/morikuni/chat/src/domain/model/aggregate"
)

type Account interface {
	GenerateID(ctx context.Context) (model.UserID, error)
	Save(ctx context.Context, account *aggregate.Account) error
	Find(ctx context.Context, id model.UserID) (*aggregate.Account, error)
	FindByEmail(ctx context.Context, email model.Email) (*aggregate.Account, error)
}
