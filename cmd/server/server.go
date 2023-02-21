package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sanpezlo/chess/internal/app"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func Start(ctx context.Context) {
	app := fx.New(
		fx.NopLogger,
		app.Module,
	)

	if err := app.Start(ctx); err != nil {
		zap.L().Fatal("Failed to start application", zap.Error(err))
	}

	<-ctx.Done()

	ctx, cf := context.WithTimeout(context.Background(), time.Second*30)
	defer cf()

	if err := app.Stop(ctx); err != nil {
		zap.L().Fatal("Failed to stop application", zap.Error(err))
	}
}

func main() {
	ctx, cf := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cf()

	Start(ctx)
}
