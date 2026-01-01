package app

import (
	"io/fs"

	"github.com/benbjohnson/clock"
	"github.com/leshless/golibrary/graceful"
	"github.com/leshless/golibrary/interrupt"
	"github.com/leshless/pet/cub/internal/config"
	"github.com/leshless/pet/cub/internal/environment"
	"github.com/leshless/pet/cub/internal/telemetry"
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
}

// @PublicValueInstance
type Adapters struct{}

// @PublicValueInstance
type Usecases struct{}

// @PublicValueInstance
type Actions struct{}

// @PublicValueInstance
type Handlers struct{}

// @PublicValueInstance
type Ports struct{}
