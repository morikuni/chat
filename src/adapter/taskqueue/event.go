package taskqueue

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/morikuni/chat/src/application/usecase"
	"github.com/morikuni/chat/src/domain/event"
	"github.com/morikuni/chat/src/infra"
	"github.com/morikuni/serializer"
	"github.com/pkg/errors"
	"google.golang.org/appengine"
	"google.golang.org/appengine/taskqueue"
)

const (
	EventHandlerPath = "/internal/event"
)

func newSerializer() serializer.Serializer {
	s := serializer.NewSerializer()
	s.Register(
		event.AccountCreated{},
	)
	return s
}

func NewEventPublisher() event.Publisher {
	return eventPublisher{
		serializer.NewByteSerializer(newSerializer()),
	}
}

type eventPublisher struct {
	serializer serializer.ByteSerializer
}

func (ep eventPublisher) Publish(c context.Context, e event.Event) error {
	task, err := ep.createTask(e)
	if err != nil {
		return err
	}
	if _, err := taskqueue.Add(c, task, ""); err != nil {
		return errors.Wrap(err, "failed to add task")
	}
	return nil
}

func (ep eventPublisher) createTask(e event.Event) (*taskqueue.Task, error) {
	data, err := ep.serializer.SerializeByte(e)
	if err != nil {
		return nil, err
	}
	h := make(http.Header)
	h.Set("Content-Type", "application/json")
	return &taskqueue.Task{
		Path:    EventHandlerPath,
		Payload: data,
		Header:  h,
		Method:  http.MethodPost,
	}, nil
}

func NewTaskHandler(
	log infra.Logger,
	eventHandler usecase.EventHandler,
) http.Handler {
	return taskHandler{
		log,
		eventHandler,
		newSerializer(),
	}
}

type taskHandler struct {
	log          infra.Logger
	eventHandler usecase.EventHandler
	serializer   serializer.Serializer
}

func (th taskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	v, err := th.serializer.Deserialize(r.Body)
	if err != nil {
		th.log.Errorf(ctx, "failed to deserialize event: %v", err)
		return
	}
	e, ok := v.(event.Event)
	if !ok {
		th.log.Errorf(ctx, "invalid domain event: %#v", v)
		return
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
