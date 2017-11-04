package event

import (
	"context"
)

type Publisher interface {
	Publish(c context.Context, event Event) error
}
