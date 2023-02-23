package users

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sanpezlo/chess/internal/db"
	"github.com/sanpezlo/chess/internal/web"
)

func (c *controller) get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := c.repository.GetUser(r.Context(), id, true)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			web.StatusNotFound(w, err)
		} else {
			web.StatusInternalServerError(w, err)
		}
		return
	}

	web.Write(w, user)
}
