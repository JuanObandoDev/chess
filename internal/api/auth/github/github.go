package github

import (
	"github.com/sanpezlo/chess/internal/services/auth"
	"go.uber.org/fx"
)

type Controller struct {
	gs *auth.GitHubService
	as *auth.Service
}

func New(gs *auth.GitHubService, as *auth.Service) *Controller {
	return &Controller{gs: gs, as: as}
}

var Module = fx.Options(
	fx.Provide(New),
)
