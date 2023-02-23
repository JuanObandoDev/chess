package web

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Error struct {
	Message    string `json:"message,omitempty"`
	Suggestion string `json:"suggested,omitempty"`
	Error      string `json:"error,omitempty"`
}

func StatusNotFound(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusNotFound)
	if err == nil {
		errToWriter(w, errors.New("Not found"))
	} else {
		errToWriter(w, err)
	}
}

func StatusInternalServerError(w http.ResponseWriter, err error) {
	if errors.Is(err, context.Canceled) {
		return
	}
	zap.L().Error("Internal error", zap.Error(err))
	w.WriteHeader(http.StatusInternalServerError)
	errToWriter(w, errors.New("Something went wrong but the details have been omitted from this error for security reasons"))
}

func StatusUnauthorized(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusUnauthorized)
	errToWriter(w, err)
}

func StatusNotAcceptable(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusNotAcceptable)
	errToWriter(w, err)
}

func StatusBadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	errToWriter(w, err)
}

type HumanReadable struct {
	err     error
	desc    string
	suggest string
}

func (h *HumanReadable) Error() string {
	return h.err.Error()
}

func WithDescription(err error, desc string) error {
	return &HumanReadable{err, desc, ""}
}

func WithSuggestion(err error, desc, suggest string) error {
	return &HumanReadable{err, desc, suggest}
}

func errToWriter(w http.ResponseWriter, err error) {
	if err == nil {
		err = errors.New("Unknown or unspecified error")
	}

	var message string
	var suggest string
	herr, ok := err.(*HumanReadable)
	if ok {
		message = herr.desc
		suggest = herr.suggest
	}

	Write(w, Error{
		Message:    message,
		Suggestion: suggest,
		Error:      err.Error(),
	})
}
