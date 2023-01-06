package adder

import (
	"context"
	"min/adder/pkg/api"
	"pkg/api"
)

type GRPCServer struct{}

func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*AddResponse, error) {
	return &api.AddResponse{}
}
