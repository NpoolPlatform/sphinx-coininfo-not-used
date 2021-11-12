package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/sphinx-coininfo/message/npool"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/core"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// https://github.com/grpc/grpc-go/issues/3794
// require_unimplemented_servers=false
type Server struct {
	npool.UnimplementedSphinxCoininfoServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterSphinxCoininfoServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterSphinxCoininfoHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}

// coininfo part

func (Server) GetCoinInfos(ctx context.Context, in *npool.GetCoinInfosRequest) (resp *npool.GetCoinInfosResponse, err error) {
	return core.GetCoinInfos(ctx, in)
}

func (Server) RegisterCoin(ctx context.Context, in *npool.RegisterCoinRequest) (resp *npool.RegisterCoinResponse, err error) {
	resp, err = core.RegisterCoin(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("register coin error: %w", err)
		return &npool.RegisterCoinResponse{Info: "failed"}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
