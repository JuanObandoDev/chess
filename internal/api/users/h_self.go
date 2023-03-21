package users

import (
	"net/http"

	"github.com/sanpezlo/chess/internal/services/auth"
	"github.com/sanpezlo/chess/internal/web"
)

func (c *controller) self(w http.ResponseWriter, r *http.Request) {
	ai, ok := auth.GetAuthenticationInfo(w, r)
	if !ok {
		return
	}

	user, err := c.repository.GetUser(r.Context(), ai.Cookie.UserID, false)
	if err != nil {
		web.StatusInternalServerError(w, err)
		return
	}

	web.Write(w, user)
}
