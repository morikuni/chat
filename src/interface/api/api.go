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
	"google.golang.org/appengine"
)

func NewAPI(
	posting usecase.Posting,
	authentication usecase.Authentication,
	chatReader reader.Chat,
	logger infra.Logger,
) API {
	return API{
		posting,
		authentication,
		chatReader,
		logger,
	}
}

type API struct {
	posting        usecase.Posting
	authentication usecase.Authentication
	chatReader     reader.Chat
	logger         infra.Logger
}

func (a API) GetChats(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	chats, err := a.chatReader.Chats(ctx)
	if err != nil {
		a.HandleError(ctx, w, err)
		return
	}
	a.JSON(ctx, w, http.StatusOK, chats)
}

func (a API) PostChats(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	message := r.FormValue("message")
	if err := a.posting.PostChat(ctx, message); err != nil {
		a.HandleError(ctx, w, err)
		return
	}
}

func (a API) PostAccounts(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	email := r.FormValue("email")
	password := r.FormValue("password")
	if err := a.authentication.CreateAccount(ctx, email, password); err != nil {
		a.HandleError(ctx, w, err)
		return
	}
}

func (a API) HandleError(ctx context.Context, w http.ResponseWriter, err error) {
	switch t := err.(type) {
	case usecase.ValidationError:
		a.JSON(ctx, w, http.StatusBadRequest, ValidationError(t))
	default:
		a.logger.Debugf(ctx, "hoge %#v", t)
		buf := &bytes.Buffer{}
		fmt.Fprintf(buf, "api: %v\n", err)
		if s, ok := err.(infra.StackTraceError); ok {
			for _, f := range s.StackTrace() {
				fmt.Fprintf(buf, "%+s:%d\n", f, f)
			}
		}
		a.logger.Errorf(ctx, "%s", buf.String())
		a.JSON(ctx, w, http.StatusInternalServerError, InternalServerError)
	}
}

func (a API) JSON(ctx context.Context, w http.ResponseWriter, status int, value interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(value)
	if err != nil {
		a.logger.Errorf(ctx, "failed to encode json: %v", err)
	}
}
