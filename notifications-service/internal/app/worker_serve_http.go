package app

import (
	"context"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/zap"
)

// serveHTTP serves HTTP.
func serveHTTP(ctx context.Context, app *App) {
	router := fiber.New(fiber.Config{
		Prefork:                  false,
		ReadTimeout:              app.cfg.Delivery.HTTPServer.ReadTimeout,
		WriteTimeout:             app.cfg.Delivery.HTTPServer.WriteTimeout,
		Network:                  fiber.NetworkTCP4,
		BodyLimit:                app.cfg.Delivery.HTTPServer.BodySizeLimitBytes,
		AppName:                  app.meta.Info.Name,
		EnableTrustedProxyCheck:  true,
		EnableSplittingOnParsers: true,
	})

	app.registerHTTPRoutes(router)

	// Middlewares
	router.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed}))
	router.Use(requestid.New())
	router.Use(recover.New())
	router.Use(favicon.New())

	go func() {
		<-ctx.Done()
		app.logger.Info("HTTP server: initiating graceful shutdown")

		sdCtx, cancel := context.WithTimeout(context.Background(), app.cfg.Delivery.HTTPServer.GracefulTimeout)
		defer cancel()

		if err := router.Shutdown(); err != nil {
			app.logger.Error("HTTP server shutdown error", zap.Error(err))
		}

		<-sdCtx.Done()
		app.logger.Info("HTTP server stopped gracefully")
	}()

	err := router.Listen(app.cfg.Delivery.HTTPServer.ListenAddress)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		app.logger.Fatal("failed to start HTTP server", zap.Error(err))
	}
}
