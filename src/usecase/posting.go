package usecase

import (
	"context"
)

type Posting interface {
	PostChat(ctx context.Context, message string) error
}
