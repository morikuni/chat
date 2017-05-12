package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/morikuni/chat/chat/usecase"
	"github.com/morikuni/yacm"
	"github.com/pkg/errors"
)

type API interface {
	yacm.Service
	Path() string
}

func Shutter(w http.ResponseWriter, r *http.Request, err error) {
	log.Println(err)
	Error(w, 500, http.StatusText(http.StatusInternalServerError))
}

func Catcher(w http.ResponseWriter, r *http.Request, err error) error {
	switch e := errors.Cause(err).(type) {
	case usecase.ValidationError:
		Error(w, 400, e.Error())
		return nil
	default:
		return err
	}
}

func Error(w http.ResponseWriter, statusCode int, message string) {
	JSON(w, statusCode, map[string]string{
		"error": message,
	})
}

func JSON(w http.ResponseWriter, statusCode int, obj interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(obj)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to encode json"))
	}
}
