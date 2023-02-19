package app

import (
	"context"
	"net"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sanpezlo/chess/internal/config"
	"github.com/sanpezlo/chess/internal/logger"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewRouter() chi.Router {
	router := chi.NewRouter()
	return router
}

func NewServer(lc fx.Lifecycle, cfg *config.Config, l *zap.Logger, router chi.Router) *http.Server {
	server := &http.Server{
		Handler: router,
		Addr:    cfg.ListenAddr,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			server.BaseContext = func(net.Listener) context.Context { return ctx }
			return nil
		},
		OnStop: func(ctx context.Context) error {
			l.Debug("HTTP server stopping")
			return server.Shutdown(ctx)
		},
	})

	return server
}

var Module = fx.Options(
	config.Module,
	logger.Module,
	fx.Provide(NewRouter),
	fx.Provide(NewServer),
	fx.Invoke(func(l *zap.Logger, server *http.Server) {
		l.Debug("HTTP server starting")
		go func() {
			if err := server.ListenAndServe(); err != http.ErrServerClosed {
				l.Fatal("HTTP server failed", zap.Error(err))
			}
		}()
	}),
)
