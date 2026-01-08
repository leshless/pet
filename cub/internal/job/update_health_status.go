package job

import (
	"context"

	healthlogic "github.com/leshless/pet/cub/internal/logic/health"
	"github.com/leshless/pet/cub/internal/telemetry"
)

const UpdateHealthStatusJobName = "update_health_status"

type UpdateHealthStatus Job

// @PublicPointerInstance
type updateHealthStatus struct {
	telemetry.Telemetry
	controller healthlogic.Controller
}

var _ UpdateHealthStatus = (*updateHealthStatus)(nil)

func (j *updateHealthStatus) Exec(ctx context.Context) error {
	_, err := j.controller.UpdateStatus(ctx, healthlogic.NewUpdateStatusArg())
	return err
}
