package auth

import (
	"context"
	"os/user"

	"github.com/go-chi/chi"
	"github.com/sanpezlo/chess/internal/api/auth/github"
	"go.uber.org/fx"
)

type OAuthProvider interface {
	Link() string
	Login(ctx context.Context, state, code string) (*user.User, error)
}

var Module = fx.Options(
	github.Module,
	fx.Invoke(func(r chi.Router, gc *github.Controller) {
		rtr := chi.NewRouter()

		r.Mount("/auth", rtr)

		rtr.Get("/github/link", gc.Link)

		rtr.Post("/github/callback", gc.Callback)

	}),
)
