package users

import (
	"net/http"

	"github.com/sanpezlo/chess/internal/web"
)

type queryParams struct {
	Sort string `qstring:"sort" valid:"in(desc|asc),optional"`
	Max  int    `qstring:"max"  valid:"min(1),optional"`
	Skip int    `qstring:"skip" valid:"min(0),optional"`
}

func (c *controller) list(w http.ResponseWriter, r *http.Request) {
	var p queryParams
	if !web.ParseQuery(w, r, &p) {
		return
	}

	if p.Sort == "" {
		p.Sort = "desc"
	}
	if p.Max == 0 {
		p.Max = 50
	} else if p.Max > 100 {
		p.Max = 100
	}

	users, err := c.repository.GetUsers(r.Context(), p.Sort, p.Max, p.Skip, true)
	if err != nil {
		web.StatusInternalServerError(w, err)
		return
	}

	web.Write(w, users)
}
