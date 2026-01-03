package grpc

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/leshless/golibrary/graceful"
	"github.com/leshless/pet/cub/internal/config"
	"github.com/leshless/pet/cub/internal/environment"
	"github.com/leshless/pet/cub/internal/telemetry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

func InitPort(
	logger telemetry.Logger,
	tel telemetry.Telemetry,
	configHolder config.Holder,
	environmentHolder environment.Holder,
	gracefulRegistrator graceful.Registrator,
	healthHanlder healthpb.HealthServer,
) (*port, error) {
	logger.Info("initializing grpc port")

	config := configHolder.Config().GRPC
	environment := environmentHolder.Environment()

	options := []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     config.KeepaliveMaxConnectionIdle,
			MaxConnectionAge:      config.KeepaliveMaxConnectionAge,
			MaxConnectionAgeGrace: config.KeepaliveMaxConnectionAgeGrace,
			Time:                  config.KeepaliveTime,
			Timeout:               config.KeepaliveTimeout,
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             config.KeepaliveMinTime,
			PermitWithoutStream: config.KeepalivePermitWithoutStream,
		}),
		grpc.MaxRecvMsgSize(config.MaxReceiveMessageSizeMB << 20),
		grpc.MaxSendMsgSize(config.MaxSendMessageSizeMB << 20),
		grpc.ConnectionTimeout(config.Timeout),
	}

	if config.EnableTLS {
		cert, err := tls.X509KeyPair([]byte(environment.TLSCertificate), []byte(environment.TLSKey))
		if err != nil {
			logger.Error("failed to create x509 key pair for GRPC TLS", telemetry.Error(err))
			return nil, fmt.Errorf("creating x509 key pair for TLS: %w", err)
		}

		creds := credentials.NewTLS(&tls.Config{
			Certificates: []tls.Certificate{cert},
			MinVersion:   tls.VersionTLS12,
			ClientAuth:   tls.NoClientCert,
		})

		options = append(options, grpc.Creds(creds))
	}

	options = append(options, grpc.ChainUnaryInterceptor(
		telemetryMiddleware(tel),
	))

	grpcServer := grpc.NewServer(options...)

	if config.EnableHealth {
		healthpb.RegisterHealthServer(grpcServer, healthHanlder)
	}

	if config.EnableReflection {
		reflection.Register(grpcServer)
	}

	gracefulRegistrator.Register(func(_ context.Context) error {
		grpcServer.Stop()
		return nil
	})

	logger.Info("grpc port successfully initialized")

	return NewPort(
		tel,
		configHolder,
		grpcServer,
	), nil
}
