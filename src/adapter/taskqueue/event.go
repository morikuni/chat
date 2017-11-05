package taskqueue

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/morikuni/chat/src/domain/event"
	"github.com/morikuni/chat/src/infra"
	"github.com/morikuni/chat/src/usecase"
	"github.com/pkg/errors"
	"google.golang.org/appengine"
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

func deserialize(se serializedEvent) (event.Event, error) {
	switch se.Name {
	case "account_created":
		var e event.AccountCreated
		if err := json.Unmarshal(se.Payload, &e); err != nil {
			return nil, errors.Wrap(err, "failed to decode json")
		}
		return e, nil
	default:
		return nil, errors.New("unknown event")
	}
}

func NewTaskHandler(
	log infra.Logger,
	eventHandler usecase.EventHandler,
) http.Handler {
	return taskHandler{
		log,
		eventHandler,
	}
}

type taskHandler struct {
	log          infra.Logger
	eventHandler usecase.EventHandler
}

func (th taskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	name := r.FormValue("name")
	payload, err := base64.StdEncoding.DecodeString(r.FormValue("payload"))
	if err != nil {
		th.log.Errorf(ctx, "failed to decode payload: %v", err)
	}
	se := serializedEvent{
		name,
		payload,
	}

	e, err := deserialize(se)
	if err != nil {
		th.log.Errorf(ctx, "failed to deserialize payload: %v", err)
	}

	if err := th.eventHandler.Handle(ctx, e); err != nil {
		buf := &bytes.Buffer{}
		fmt.Fprintf(buf, "taskqueue: %v\n", err)
		if s, ok := err.(infra.StackTraceError); ok {
			for _, f := range s.StackTrace() {
				fmt.Fprintf(buf, "%+s:%d\n", f, f)
			}
		}
		th.log.Errorf(ctx, "%s", buf.String())
	}
}
