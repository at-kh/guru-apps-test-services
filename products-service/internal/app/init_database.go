package app

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
	"go.uber.org/zap"
)

// initDatabase initializes database connection and migrations.
func (a *App) initDatabase() {
	var (
		db  *sqlx.DB
		err error
	)

	for attempt := 1; attempt <= a.cfg.Storage.Postgres.MaxRetries; attempt++ {
		db, err = a.connectDB()
		if err == nil {
			break
		}

		a.logger.Warn("db connection failed",
			zap.Int("attempt", attempt),
			zap.Int("max_retries", a.cfg.Storage.Postgres.MaxRetries),
			zap.Error(err),
		)

		time.Sleep(a.cfg.Storage.Postgres.RetryDelay)
	}

	if err != nil {
		a.logger.Fatal("couldn't connect to db after retries", zap.Error(err))
	}

	if a.cfg.Storage.Postgres.AutoMigrate {
		a.logger.Info("auto-migration enabled")

		for attempt := 1; attempt <= a.cfg.Storage.Postgres.MaxRetries; attempt++ {
			if err = a.runMigrationsWithDB(db); err == nil {
				break
			}

			a.logger.Warn("migration failed, retrying",
				zap.Int("attempt", attempt),
				zap.Int("max_retries", a.cfg.Storage.Postgres.MaxRetries),
				zap.Error(err),
			)
			time.Sleep(a.cfg.Storage.Postgres.RetryDelay)
		}

		if err != nil {
			a.logger.Fatal("failed to apply migrations after retries", zap.Error(err))
		}
	}

	a.db = db
}

// connectDB creates and returns a configured DB connection.
func (a *App) connectDB() (*sqlx.DB, error) {
	db, err := sqlx.Open(a.cfg.Storage.Postgres.Driver, a.cfg.Storage.Postgres.DSN)
	if err != nil {
		return nil, fmt.Errorf("sql.Open error: %w", err)
	}

	if a.cfg.Storage.Postgres.ConnMaxOpenNum > 0 {
		db.SetMaxOpenConns(a.cfg.Storage.Postgres.ConnMaxOpenNum)
	}
	if a.cfg.Storage.Postgres.ConnMaxIdleNum > 0 {
		db.SetMaxIdleConns(a.cfg.Storage.Postgres.ConnMaxIdleNum)
	}
	if a.cfg.Storage.Postgres.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(a.cfg.Storage.Postgres.ConnMaxLifetime)
	}

	ctx := context.Background()
	if a.cfg.Storage.Postgres.QueryTimeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, a.cfg.Storage.Postgres.QueryTimeout)
		defer cancel()
	}

	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("db ping failed: %w", err)
	}

	return db, nil
}

// runMigrationsWithDB applies database migrations from embed.FS
func (a *App) runMigrationsWithDB(db *sqlx.DB) error {
	migrations := &migrate.AssetMigrationSource{
		Asset: func(name string) ([]byte, error) {
			return a.dbMigrationsFS.ReadFile(name)
		},
		AssetDir: func(name string) ([]string, error) {
			entries, err := a.dbMigrationsFS.ReadDir(name)
			if err != nil {
				return nil, err
			}
			names := make([]string, 0, len(entries))
			for _, e := range entries {
				names = append(names, e.Name())
			}
			return names, nil
		},
		Dir: a.cfg.Storage.Postgres.MigrationDirectory,
	}

	direction := migrate.Up
	if strings.ToLower(a.cfg.Storage.Postgres.MigrationDirection) == "down" {
		direction = migrate.Down
	}

	applied, err := migrate.Exec(db.DB, a.cfg.Storage.Postgres.Dialect, migrations, direction)
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	if applied > 0 {
		a.logger.Info("migrations applied", zap.Int("count", applied))
	} else {
		a.logger.Info("no migrations to apply")
	}

	return nil
}
