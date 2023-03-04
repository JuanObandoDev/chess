package web

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/dyninc/qstring"
	"github.com/gorilla/schema"
	"github.com/pkg/errors"
)

func ParseQuery(w http.ResponseWriter, r *http.Request, out interface{}) bool {
	if err := qstring.Unmarshal(r.URL.Query(), out); err != nil {
		StatusBadRequest(w, err)
		return false
	}
	if _, err := govalidator.ValidateStruct(out); err != nil {
		StatusBadRequest(w, err)
		return false
	}
	return true
}

func ParseBody(w http.ResponseWriter, r *http.Request, out interface{}) bool {
	if err := json.NewDecoder(r.Body).Decode(out); err != nil {
		StatusBadRequest(w, WithSuggestion(err,
			"Could not process request data",
			"Please try again, if the issue persists contact the support team."))
		return false
	}
	if _, err := govalidator.ValidateStruct(out); err != nil {
		StatusBadRequest(w, WithSuggestion(err,
			"Could not validate request",
			"Please try again, if the issue persists contact the support team."))
		return false
	}
	return true
}

func DecodeBody(r *http.Request, v interface{}) error {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		return json.NewDecoder(r.Body).Decode(v)

	case "application/x-www-form-urlencoded":
		err := r.ParseForm()
		if err != nil {
			return errors.Wrap(err, "failed to parse form")
		}
		return schema.NewDecoder().Decode(v, r.PostForm)
	}
	return errors.Errorf("cannot decode content type %s", r.Header.Get("Content-Type"))
}
