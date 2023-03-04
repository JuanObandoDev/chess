package web

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

func Write(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)

	err := encoder.Encode(data)

	if err != nil {
		StatusInternalServerError(w, errors.Wrap(err, "Failed to write payload"))
		return
	}
}
