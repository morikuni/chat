package datastore

import (
	"context"

	"github.com/morikuni/chat/src/infra"
	oldcontext "golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

func NewTransactionManager() infra.TransactionManager {
	return transaction{}
}

type transaction struct{}

func (transaction) Exec(ctx context.Context, f func(ctx context.Context) error) error {
	if ctx.Value(txKey) != nil {
		return f(ctx)
	}
	ctx = context.WithValue(ctx, txKey, struct{}{})

	f2 := func(ctx oldcontext.Context) error {
		return f(ctx)
	}
	return datastore.RunInTransaction(ctx, f2, &datastore.TransactionOptions{XG: true})
}

type tx struct{}

var txKey tx
