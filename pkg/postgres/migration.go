package postgres

import (
	"bookwormia/pkg/logger"
	"context"
	"database/sql"
	"os"
	"path/filepath"
)

func Migrate(ctx context.Context, path string, db *sql.DB) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			migrationSQL, err := os.ReadFile(filepath.Join(path, file.Name()))

			if err != nil {
				logger.Logger.Error().Err(err).Msg("error while reading migration files")
				return err
			}

			_, err = db.ExecContext(context.Background(), string(migrationSQL))
			if err != nil {
				logger.Logger.Error().Err(err).Msg("error while applying new migrations")
				return err
			}
		}
	}

	return nil
}
