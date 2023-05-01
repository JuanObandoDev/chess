package games

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sanpezlo/chess/internal/services/games"
	"github.com/sanpezlo/chess/internal/web"
	"go.uber.org/fx"
)

type controller struct {
	gs *games.Service
}

func New(gs *games.Service) *controller {
	return &controller{gs: gs}
}

var Module = fx.Options(
	fx.Provide(New),
	fx.Invoke(func(r chi.Router, c *controller) {
		rtr := chi.NewRouter()
		r.Mount("/games", rtr)

		rtr.Mount("/ws", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := c.gs.Websocket.HandleRequest(w, r)
			if err != nil {
				web.StatusInternalServerError(w, err)
			}
		}))
	}),
)
