package services

import (
	"github.com/sanpezlo/chess/internal/services/auth"
	"go.uber.org/fx"
)

var Module = fx.Options(
	auth.Module,
)
