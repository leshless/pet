package grpc

import (
	"context"

	"github.com/leshless/pet/cub/internal/telemetry"
	"google.golang.org/grpc"
)

func telemetryMiddleware(tel telemetry.Telemetry) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		logger := tel.Logger.With(telemetry.Call(info.FullMethod))

		tel.Logger.Info("recieved unary gRPC call")

		response, err := handler(ctx, req)

		if err != nil {
			logger.Error("processed GRPC call with error", telemetry.Error(err))
		} else {
			logger.Info("sucessfully processed unary GRPC call")
		}

		return response, err
	}
}
