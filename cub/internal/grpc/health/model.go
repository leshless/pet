package health

import (
	healthmodel "github.com/leshless/pet/cub/internal/model/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

var StatusFromModel = map[healthmodel.Status]healthpb.HealthCheckResponse_ServingStatus{
	healthmodel.StatusUnknown:    healthpb.HealthCheckResponse_SERVICE_UNKNOWN,
	healthmodel.StatusServing:    healthpb.HealthCheckResponse_SERVING,
	healthmodel.StatusNotServing: healthpb.HealthCheckResponse_NOT_SERVING,
}
