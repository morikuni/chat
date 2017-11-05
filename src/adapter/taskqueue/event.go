package taskqueue

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/url"

	"github.com/morikuni/chat/src/domain/event"
	"github.com/pkg/errors"
	"google.golang.org/appengine/taskqueue"
)

const (
	EventHandlerPath = "/internal/event"
)

func NewEventPublisher() event.Publisher {
	return eventPublisher{}
}

type eventPublisher struct{}

func (ep eventPublisher) Publish(c context.Context, e event.Event) error {
	task, err := createTask(e)
	if err != nil {
		return err
	}
	if _, err := taskqueue.Add(c, task, ""); err != nil {
		return errors.Wrap(err, "failed to add task")
	}
	return nil
}

func createTask(e event.Event) (*taskqueue.Task, error) {
	se, err := serialize(e)
	if err != nil {
		return nil, err
	}
	v := url.Values{}
	v.Set("name", se.Name)
	v.Set("payload", base64.StdEncoding.EncodeToString(se.Payload))
	return taskqueue.NewPOSTTask(EventHandlerPath, v), nil
}

type serializedEvent struct {
	Name    string
	Payload []byte
}

func serialize(e event.Event) (serializedEvent, error) {
	var name string
	switch e.(type) {
	case event.AccountCreated:
		name = "account_created"
	default:
		return serializedEvent{}, errors.Errorf("unknown event: %#v", e)
	}
	payload, err := json.Marshal(e)
	if err != nil {
		return serializedEvent{}, errors.Wrap(err, "failed to encode json")
	}
	return serializedEvent{
		name,
		payload,
	}, nil
}
