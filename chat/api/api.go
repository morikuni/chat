package api

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

func JSON(w http.ResponseWriter, obj interface{}) error {
	w.Header().Set("content-type", "application/json")
	err := json.NewEncoder(w).Encode(obj)
	if err != nil {
		return errors.Wrap(err, "failed to encode json")
	}
	return nil
}
