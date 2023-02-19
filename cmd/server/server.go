package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/sanpezlo/chess/internal/app"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.NopLogger,
		app.Module,
	).Run()

}
