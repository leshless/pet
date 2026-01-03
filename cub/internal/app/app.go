package app

import (
	"fmt"

	"github.com/leshless/pet/cub/internal/telemetry"
	"golang.org/x/sync/errgroup"
)

// @PublicValueInstance
type App struct {
	Primitives
	Dependencies
	Adapters
	Usecases
	Actions
	Handlers
	Ports
}

func (app *App) Run() error {
	app.Logger.Info("starting app")

	ctx := app.Interrupter.Context()

	var eg errgroup.Group

	eg.Go(func() error {
		return app.GRPC.Run(ctx)
	})

	err := eg.Wait()
	if err != nil {
		app.Logger.Error("app stopped with error", telemetry.Error(err))
		return fmt.Errorf("stopping app: %w", err)
	}

	app.Logger.Info("app successfully stopped")

	return nil
}
