package grpc

import (
	"context"

	"google.golang.org/grpc"
)

type Port interface {
	Run(ctx context.Context) error
}

// @PublicPointerInstance
type port struct {
	grpcServer *grpc.Server
}

var _ Port = (*port)(nil)

func (p *port) Run(ctx context.Context) error {
	return nil
}
