package db

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/leshless/pet/cub/internal/model"
)

func migrationFromModel(migration model.Migration) Migration {
	return Migration{
		Version:   int32(migration.Version),
		Name:      migration.Name,
		Query:     migration.Query,
		AppliedAt: pgtype.Timestamptz{Time: migration.AppliedAt, Valid: true},
	}
}

func migrationToModel(migration Migration) model.Migration {
	return model.NewMigration(
		uint(migration.Version),
		migration.Name,
		migration.Query,
		migration.AppliedAt.Time,
	)
}
