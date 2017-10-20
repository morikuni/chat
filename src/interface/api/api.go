package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/morikuni/chat/src/infra"
	"github.com/morikuni/chat/src/reader"
	"github.com/morikuni/chat/src/usecase"
	"github.com/pkg/errors"
	"google.golang.org/appengine"
)

func NewAPI(
	posting usecase.Posting,
	chatReader reader.Chat,
	logger infra.Logger,
) API {
	return API{
		posting,
		chatReader,
		logger,
	}
}

type API struct {
	posting    usecase.Posting
	chatReader reader.Chat
	log        infra.Logger
}

func (a API) GetChats(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	chats, err := a.chatReader.Chats(ctx)
	if err != nil {
		a.HandleError(ctx, w, err)
		return
	}
	a.JSON(ctx, w, chats)
}

func (a API) PostChats(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	message := r.FormValue("message")
	if err := a.posting.PostChat(ctx, message); err != nil {
		a.HandleError(ctx, w, err)
		return
	}
}

func (a API) HandleError(ctx context.Context, w http.ResponseWriter, err error) {
	switch t := errors.Cause(err).(type) {
	case usecase.ValidationError:
		w.WriteHeader(http.StatusBadRequest)
		a.JSON(ctx, w, ValidationError(t))
	default:
		buf := &bytes.Buffer{}
		fmt.Fprintf(buf, "api: %v\n", err)
		if s, ok := err.(infra.StackTraceError); ok {
			for _, f := range s.StackTrace() {
				fmt.Fprintf(buf, "%+s:%d\n", f, f)
			}
		}
		a.log.Errorf(ctx, "%s", buf.String())
		w.WriteHeader(http.StatusInternalServerError)
		a.JSON(ctx, w, InternalServerError)
	}
}

func (a API) JSON(ctx context.Context, w http.ResponseWriter, value interface{}) {
	err := json.NewEncoder(w).Encode(value)
	if err != nil {
		a.log.Errorf(ctx, "failed to encode json: %v", err)
	}
}
