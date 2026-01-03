package health

import (
	"context"

	healthlogic "github.com/leshless/pet/cub/internal/logic/health"
	"github.com/leshless/pet/cub/internal/telemetry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

// @PublicPointerInstance
type handler struct {
	telemetry.Telemetry
	checkUseCase healthlogic.CheckUseCase
}

var _ healthpb.HealthServer = (*handler)(nil)

func (h *handler) Check(ctx context.Context, req *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	arg := CheckRequestToDTO(req)

	res, err := h.checkUseCase.Exec(ctx, arg)
	if err != nil {
		return nil, err
	}

	return CheckResponseFromDTO(res), nil
}

func (h *handler) List(ctx context.Context, req *healthpb.HealthListRequest) (*healthpb.HealthListResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (h *handler) Watch(req *healthpb.HealthCheckRequest, srv grpc.ServerStreamingServer[healthpb.HealthCheckResponse]) error {
	return status.Error(codes.Unimplemented, "")
}
