package health

import (
	"context"

	healthmodel "github.com/leshless/pet/cub/internal/model/health"
	"github.com/leshless/pet/cub/internal/telemetry"
	"github.com/leshless/pet/library/go/usecase"
)

// @PublicValueInstance
type CheckArg struct {
	Service string
}

// @PublicValueInstance
type CheckRes struct {
	Status healthmodel.Status
}

type CheckUseCase usecase.UseCase[CheckArg, CheckRes]

// @PublicPointerInstance
type checkUseCase struct {
	telemetry.Telemetry
}

var _ CheckUseCase = (*checkUseCase)(nil)

func (uc *checkUseCase) Exec(ctx context.Context, arg CheckArg) (CheckRes, error) {
	return CheckRes{
		Status: healthmodel.StatusServing,
	}, nil
}
