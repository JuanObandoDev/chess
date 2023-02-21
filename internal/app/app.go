package app

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/sanpezlo/chess/internal/config"
	"github.com/sanpezlo/chess/internal/db"
	"github.com/sanpezlo/chess/internal/logger"
	"github.com/sanpezlo/chess/internal/version"
	"github.com/sanpezlo/chess/internal/web"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Options(
	config.Module,
	logger.Module,
	db.Module,
	fx.Provide(New),
	fx.Invoke(func(router chi.Router) {
		router.Get("/version", func(w http.ResponseWriter, r *http.Request) {
			web.Write(w, map[string]string{"version": version.Version})
		})
	}),
)

func New(lc fx.Lifecycle, cfg *config.Config, l *zap.Logger) chi.Router {
	router := chi.NewRouter()

	origins := []string{
		cfg.PublicWebAddress,
	}

	l.Debug("Preparing router", zap.Strings("origins", origins))

	router.Use(
		web.WithLogger,
		web.WithContentType,
		cors.Handler(cors.Options{
			AllowedOrigins:   origins,
			AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "Content-Length", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link", "Content-Length", "X-Ratelimit-Limit", "X-Ratelimit-Reset"},
			AllowCredentials: true,
			MaxAge:           300,
		}),
	)

	server := &http.Server{
		Handler: router,
		Addr:    cfg.ListenAddr,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			l.Debug("HTTP server starting")
			go func() {
				if err := server.ListenAndServe(); err != http.ErrServerClosed {
					l.Fatal("HTTP server failed", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			l.Debug("HTTP server stopping")
			return server.Shutdown(ctx)
		},
	})

	return router
}
