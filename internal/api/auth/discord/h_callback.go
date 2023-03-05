package discord

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/sanpezlo/chess/internal/web"
)

type Callback struct {
	State string `json:"state"`
	Code  string `json:"code"`
}

func (c *Controller) Callback(w http.ResponseWriter, r *http.Request) {
	var payload Callback
	if err := web.DecodeBody(r, &payload); err != nil {
		web.StatusBadRequest(w, errors.Wrap(err, "failed to decode callback payload"))
		return
	}

	user, err := c.ds.Login(r.Context(), payload.State, payload.Code)
	if err != nil {
		web.StatusBadRequest(w, err)
		return
	}

	c.as.EncodeAuthCookie(w, user)
	web.Write(w, user)
}
