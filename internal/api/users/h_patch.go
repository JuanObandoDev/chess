package users

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sanpezlo/chess/internal/db"
	"github.com/sanpezlo/chess/internal/web"
)

type patchPayload struct {
	Email *string `json:"email"`
	Name  *string `json:"name"`
	Bio   *string `json:"bio"`
}

func (c *controller) patch(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var p patchPayload
	if !web.ParseBody(w, r, &p) {
		return
	}

	user, err := c.repository.UpdateUser(r.Context(), id, p.Email, p.Name, p.Bio)
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
