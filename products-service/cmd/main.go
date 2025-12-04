package main

import (
	"context"
	"embed"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/domain/health"
	"github.com/at-kh/guru-apps-test-services/products-service/internal/app"
	"github.com/at-kh/guru-apps-test-services/products-service/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
	_ "go.uber.org/automaxprocs"
)

// Block information for the application.
var (
	name    string
	commit  string
	version string
	date    string
)

//go:embed dbschema/migrations
var dbMigrationFS embed.FS

// main is the entry point for the application.
func main() {
	ctx, cancel := registerGracefulShutdown()
	defer cancel()

	cfgPath := flag.String("c", config.DefaultPath, "configuration file")
	flag.Parse()

	app.New(
		app.Meta{
			ConfigPath: *cfgPath,
			Info: health.Info{
				Name:         name,
				BuildCommit:  commit,
				BuildDate:    date,
				BuildVersion: version,
			},
		},
	).WithMigrationFS(dbMigrationFS).Run(ctx)
}

// registerGracefulShutdown returns a context that is canceled on signals.
func registerGracefulShutdown() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer signal.Stop(signals)

		sig := <-signals
		log.Printf("[signal] received=%s action=graceful-shutdown", sig)

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)

		cancel()

		select {
		case sig2 := <-signals:
			log.Printf("[signal] second received: forcing exit (%s)", sig2)
			os.Exit(1)
		case <-shutdownCtx.Done():
			shutdownCancel()
			return
		}
	}()

	return ctx, cancel
}
