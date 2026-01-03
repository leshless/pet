package app

import (
	"context"
	"fmt"
	"io/fs"

	"github.com/benbjohnson/clock"
	"github.com/leshless/golibrary/graceful"
	"github.com/leshless/golibrary/interrupt"
	"github.com/leshless/golibrary/stupid"
	"github.com/leshless/pet/cub/internal/config"
	"github.com/leshless/pet/cub/internal/environment"
	"github.com/leshless/pet/cub/internal/grpc"
	healthgrpc "github.com/leshless/pet/cub/internal/grpc/health"
	healthlogic "github.com/leshless/pet/cub/internal/logic/health"
	"github.com/leshless/pet/cub/internal/telemetry"
	"go.uber.org/dig"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func InitApp(primitives Primitives) (*App, error) {
	c := dig.New()

	providePrimitives(c, primitives)
	provideDependencies(c)
	provideClients(c)
	provideAdapters(c)
	provideUsecases(c)
	provideActions(c)
	provideHandlers(c)
	providePorts(c)

	c.Provide(NewApp)

	var app App
	err := c.Invoke(func(a App) {
		app = a
	})
	if err != nil {
		return nil, fmt.Errorf("resolving app from the DI container: %w", err)
	}

	app.Logger.Info("app successfully initialized")

	go func() {
		<-app.Interrupter.Context().Done()

		app.Logger.Warn("graceful shutdown initiated")

		err := app.GracefulTerminator.Terminate(context.TODO())
		if err != nil {
			app.Logger.Error("app terminated with error", telemetry.Error(err))
		}
	}()

	return &app, nil
}

func providePrimitives(c *dig.Container, primitives Primitives) {
	c.Provide(stupid.Reflect(primitives.Clock), dig.As(new(clock.Clock)))
	c.Provide(stupid.Reflect(primitives.Interrupter), dig.As(new(interrupt.Interrupter)))
	c.Provide(stupid.Reflect(primitives.FS), dig.As(new(fs.FS)))

	c.Provide(stupid.Reflect(primitives))
}

func provideDependencies(c *dig.Container) {
	c.Provide(graceful.NewManager, dig.As(new(graceful.Registrator), new(graceful.Terminator)))
	c.Provide(config.InitHolder, dig.As(new(config.Holder)))
	c.Provide(environment.InitHolder, dig.As(new(environment.Holder)))
	c.Provide(telemetry.InitLogger, dig.As(new(telemetry.Logger)))
	c.Provide(telemetry.NewTelemetry)

	c.Provide(NewDependencies)
}

func provideAdapters(c *dig.Container) {

	c.Provide(NewAdapters)
}

func provideClients(c *dig.Container) {

	c.Provide(NewClients)
}

func provideUsecases(c *dig.Container) {
	c.Provide(healthlogic.NewCheckUseCase, dig.As(new(healthlogic.CheckUseCase)))

	c.Provide(NewUsecases)
}

func provideActions(c *dig.Container) {

	c.Provide(NewActions)
}

func provideHandlers(c *dig.Container) {
	c.Provide(healthgrpc.NewHandler, dig.As(new(healthpb.HealthServer)))

	c.Provide(NewHandlers)
}

func providePorts(c *dig.Container) {
	c.Provide(grpc.InitPort, dig.As(new(grpc.Port)))

	c.Provide(NewPorts)
}
