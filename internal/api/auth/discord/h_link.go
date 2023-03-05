package discord

import (
	"net/http"

	"github.com/sanpezlo/chess/internal/web"
)

type Link struct {
	URL string `json:"url"`
}

func (c *Controller) Link(w http.ResponseWriter, r *http.Request) {
	web.Write(w, Link{URL: c.ds.Link()})
}
