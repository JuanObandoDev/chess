package discord

import (
	"github.com/sanpezlo/chess/internal/services/auth"
	"go.uber.org/fx"
)

type Controller struct {
	ds *auth.DiscordService
	as *auth.Service
}

func New(ds *auth.DiscordService, as *auth.Service) *Controller {
	return &Controller{ds: ds, as: as}
}

var Module = fx.Options(
	fx.Provide(New),
)
