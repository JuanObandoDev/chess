package api

import (
	"github.com/sanpezlo/chess/internal/api/auth"
	"github.com/sanpezlo/chess/internal/api/users"
	"go.uber.org/fx"
)

var Module = fx.Options(
	users.Module,
	auth.Module,
)
