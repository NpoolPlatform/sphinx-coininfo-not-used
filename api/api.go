package api

import (
	"context"

	"github.com/NpoolPlatform/sphinx-coininfo/message/npool"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/core"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (Server) GetCoinInfos(ctx context.Context, in *emptypb.Empty) (resp *npool.GetCoinInfosResponse, err error) {
	return core.GetCoinInfos(ctx, in)
}
