package db

import (
	"context"

	"github.com/leshless/golibrary/xslices"
	"github.com/leshless/pet/cub/internal/model"
	"github.com/leshless/pet/cub/internal/telemetry"
)

type MigrationAdapter interface {
	SelectAll(ctx context.Context) ([]model.Migration, error)
}

// @PublicPointerInstance
type migrationAdapter struct {
	telemetry.Telemetry
	queries *Queries
}

var _ MigrationAdapter = (*migrationAdapter)(nil)

func (a *migrationAdapter) SelectAll(ctx context.Context) ([]model.Migration, error) {
	migrations, err := a.queries.SelectAllMigrations(ctx)
	if err != nil {
		a.Logger.Error(ctx, "failed to select migration from db", telemetry.Error(err))
		return nil, err
	}

	return xslices.Map(migrations, migrationToModel), nil
}
