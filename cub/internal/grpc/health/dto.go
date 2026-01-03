package health

import (
	healthlogic "github.com/leshless/pet/cub/internal/logic/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func CheckRequestToDTO(req *healthpb.HealthCheckRequest) healthlogic.CheckArg {
	return healthlogic.NewCheckArg(req.GetService())
}

func CheckResponseFromDTO(res healthlogic.CheckRes) *healthpb.HealthCheckResponse {
	status, ok := StatusFromModel[res.Status]
	if !ok {
		status = healthpb.HealthCheckResponse_UNKNOWN
	}

	return &healthpb.HealthCheckResponse{
		Status: status,
	}
}
