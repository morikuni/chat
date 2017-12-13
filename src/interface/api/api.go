package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/morikuni/chat/src/application"
	"github.com/morikuni/chat/src/application/reader"
	"github.com/morikuni/chat/src/application/usecase"
	"github.com/morikuni/chat/src/infra"
	"google.golang.org/appengine"
)

func NewAPI(
	posting usecase.Posting,
	authentication usecase.Authentication,
	authorization usecase.Authorization,
	chatReader reader.Chat,
	logger infra.Logger,
) API {
	return API{
		posting,
		authentication,
		authorization,
		chatReader,
		logger,
	}
}

type API struct {
	posting        usecase.Posting
	authentication usecase.Authentication
	authorization  usecase.Authorization
	chatReader     reader.Chat
	logger         infra.Logger
}

func (a API) GetChats(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	cursorToken := r.FormValue("cursor_token")
	chats, err := a.chatReader.Chats(ctx, cursorToken)
	if err != nil {
		a.HandleError(ctx, w, err)
		return
	}
	a.JSON(ctx, w, http.StatusOK, chats)
}

func (a API) PostChats(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	_, err := a.authorization.ValidateAccessToken(ctx, r.FormValue("token"))
	if err != nil {
		a.HandleError(ctx, w, err)
	}

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

func (a API) PostTokens(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	email := r.FormValue("email")
	password := r.FormValue("password")
	at, err := a.authorization.GenerateAccessToken(ctx, email, password)
	if err != nil {
		a.HandleError(ctx, w, err)
		return
	}
	a.JSON(ctx, w, http.StatusOK, at)
}

func (a API) HandleError(ctx context.Context, w http.ResponseWriter, err error) {
	var (
		status int
		e      Error
	)
	switch t := err.(type) {
	case application.ValidationError:
		status, e = http.StatusBadRequest, ValidationError(t)
	case application.InvalidCredentialError:
		status, e = http.StatusForbidden, InvalidCredentialError
	default:
		status, e = http.StatusInternalServerError, InternalServerError
	}

	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "api: %T(%v)\n", err, err)
	if s, ok := err.(infra.StackTraceError); ok {
		for _, f := range s.StackTrace() {
			fmt.Fprintf(buf, "%+s:%d\n", f, f)
		}
	}

	if status/100 == 5 {
		a.logger.Errorf(ctx, "%s", buf.String())
	} else {
		a.logger.Debugf(ctx, "%s", buf.String())
	}
	a.JSON(ctx, w, status, e)
}

func (a API) JSON(ctx context.Context, w http.ResponseWriter, status int, value interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(value)
	if err != nil {
		a.logger.Errorf(ctx, "failed to encode json: %v", err)
	}
}
