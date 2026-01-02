package grpc

import (
	"github.com/leshless/golibrary/graceful"
	"github.com/leshless/pet/cub/internal/config"
	"github.com/leshless/pet/cub/internal/environment"
	"github.com/leshless/pet/cub/internal/telemetry"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func InitPort(
	logger telemetry.Logger,
	configHolder config.Holder,
	environmentHolder environment.Holder,
	gracefulRegistrator graceful.Registrator,
	healthHanlder healthpb.HealthServer,
) (*port, error) {
	grpcServer := grpc.NewServer()

	healthpb.RegisterHealthServer(grpcServer, healthHanlder)
}
