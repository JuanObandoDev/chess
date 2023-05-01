package services

import (
	"github.com/sanpezlo/chess/internal/services/auth"
	"github.com/sanpezlo/chess/internal/services/games"
	"go.uber.org/fx"
)

var Module = fx.Options(
	auth.Module,
	games.Module,
)
