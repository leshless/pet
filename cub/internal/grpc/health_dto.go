package grpc

import (
	healthlogic "github.com/leshless/pet/cub/internal/logic/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func checkRequestToDTO(req *healthpb.HealthCheckRequest) healthlogic.CheckArg {
	return healthlogic.NewCheckArg()
}

func checkResponseFromDTO(res healthlogic.CheckRes) *healthpb.HealthCheckResponse {
	status, ok := healthStatusFromModel[res.Status]
	if !ok {
		status = healthpb.HealthCheckResponse_UNKNOWN
	}

	return &healthpb.HealthCheckResponse{
		Status: status,
	}
}
