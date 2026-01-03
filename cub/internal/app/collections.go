package app

import (
	"io/fs"

	"github.com/benbjohnson/clock"
	"github.com/leshless/golibrary/graceful"
	"github.com/leshless/golibrary/interrupt"
	"github.com/leshless/pet/cub/internal/config"
	"github.com/leshless/pet/cub/internal/environment"
	"github.com/leshless/pet/cub/internal/grpc"
	healthlogic "github.com/leshless/pet/cub/internal/logic/health"
	"github.com/leshless/pet/cub/internal/telemetry"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// @PublicValueInstance
type Primitives struct {
	Clock       clock.Clock
	Interrupter interrupt.Interrupter
	FS          fs.FS
}

// @PublicValueInstance
type Dependencies struct {
	GracefulRegistrator graceful.Registrator
	GracefulTerminator  graceful.Terminator
	ConfigHolder        config.Holder
	EnvironmentHolder   environment.Holder
	Logger              telemetry.Logger
	Telemetry           telemetry.Telemetry
}

// @PublicValueInstance
type Clients struct{}

// @PublicValueInstance
type Adapters struct{}

// @PublicValueInstance
type Usecases struct {
	CheckHealth healthlogic.CheckUseCase
}

// @PublicValueInstance
type Actions struct{}

// @PublicValueInstance
type Handlers struct {
	HealthGRPC healthpb.HealthServer
}

// @PublicValueInstance
type Ports struct {
	GRPC grpc.Port
}
