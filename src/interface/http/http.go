package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/morikuni/chat/src/errors"

	"github.com/morikuni/chat/src/domain"
	"github.com/morikuni/chat/src/usecase"

	"github.com/morikuni/failure"
)

type Server struct {
	addr    string
	account usecase.Account
	message usecase.Message
	room    usecase.Room
}

func NewServer(
	addr string,
	account usecase.Account,
	message usecase.Message,
	room usecase.Room,
) *Server {
	return &Server{
		addr,
		account,
		message,
		room,
	}
}

func (s *Server) Run(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_account", s.createAccount)
	mux.HandleFunc("/create_room", s.createRoom)
	mux.HandleFunc("/post_message", s.postMessage)

	server := &http.Server{
		Addr:    s.addr,
		Handler: mux,
	}

	errChan := make(chan error, 1)
	go func() {
		err := server.ListenAndServe()
		errChan <- failure.Wrap(err)
	}()

	select {
	case <-ctx.Done():
		err := server.Shutdown(context.Background())
		if err != nil {
			return failure.Wrap(err)
		}
	case err := <-errChan:
		if err != nil && err != http.ErrServerClosed {
			return failure.Wrap(err)
		}
		return nil
	}

	err := <-errChan
	if err != nil && err != http.ErrServerClosed {
		return failure.Wrap(err)
	}

	return nil
}

func (s *Server) do(w http.ResponseWriter, r *http.Request, data interface{}, f func(ctx context.Context) (interface{}, error)) {
	type Error struct {
		Error string `json:"error"`
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		err := json.NewEncoder(w).Encode(&Error{Error: "please use POST"})
		if err != nil {
			log.Println(r.URL.Path, failure.Wrap(err))
		}
		return
	}

	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(&Error{Error: "invalid json"})
		if err != nil {
			log.Println(r.URL.Path, failure.Wrap(err))
		}
		return
	}

	res, err := f(r.Context())
	if err != nil {
		log.Println(r.URL.Path, failure.Wrap(err))

		switch c, _ := failure.CodeOf(err); c {
		case errors.InvalidArgument:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}

		msg, ok := failure.MessageOf(err)
		if !ok {
			msg = "internal error"
		}

		err := json.NewEncoder(w).Encode(&Error{Error: msg})
		if err != nil {
			log.Println(r.URL.Path, failure.Wrap(err))
		}
		return
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(&Error{Error: "internal error"})
		if err != nil {
			log.Println(r.URL.Path, failure.Wrap(err))
		}
	}
}

func (s *Server) createAccount(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Name string `json:"name"`
	}
	type Response struct {
		ID domain.AccountID `json:"id"`
	}

	var req Request
	s.do(w, r, &req, func(ctx context.Context) (interface{}, error) {
		name, err := domain.NewAccountName(req.Name)
		if err != nil {
			return nil, failure.Wrap(err)
		}

		res, err := s.account.CreateAccount(r.Context(), &usecase.CreateAccountRequest{
			Name: name,
		})
		if err != nil {
			return nil, failure.Wrap(err)
		}

		return &Response{
			ID: res.Account.ID,
		}, nil
	})
}

func (s *Server) createRoom(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		LoginUserID string `json:"login_user_id"`
		Name        string `json:"name"`
	}
	type Response struct {
		ID domain.RoomID `json:"id"`
	}

	var req Request
	s.do(w, r, &req, func(ctx context.Context) (interface{}, error) {
		loginID, err := domain.NewAccountID(req.LoginUserID)
		if err != nil {
			return nil, failure.Wrap(err)
		}

		name, err := domain.NewRoomName(req.Name)
		if err != nil {
			return nil, failure.Wrap(err)
		}

		res, err := s.room.CreateRoom(r.Context(), &usecase.CreateRoomRequest{
			OwnerID: loginID,
			Name:    name,
		})
		if err != nil {
			return nil, failure.Wrap(err)
		}

		return &Response{
			ID: res.Room.ID,
		}, nil
	})
}

func (s *Server) postMessage(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		LoginUserID string `json:"login_user_id"`
		RoomID      string `json:"room_id"`
		Message     string `json:"message"`
	}
	type Response struct {
		ID domain.MessageID `json:"id"`
	}

	var req Request
	s.do(w, r, &req, func(ctx context.Context) (interface{}, error) {
		loginID, err := domain.NewAccountID(req.LoginUserID)
		if err != nil {
			return nil, failure.Wrap(err)
		}

		roomID, err := domain.NewRoomID(req.RoomID)
		if err != nil {
			return nil, failure.Wrap(err)
		}

		body, err := domain.NewMessageBody(req.Message)
		if err != nil {
			return nil, failure.Wrap(err)
		}

		res, err := s.message.PostMessage(r.Context(), &usecase.PostMessageRequest{
			PostedBy: loginID,
			RoomID:   roomID,
			Body:     body,
		})
		if err != nil {
			return nil, failure.Wrap(err)
		}

		return &Response{
			ID: res.Message.ID,
		}, nil
	})
}
