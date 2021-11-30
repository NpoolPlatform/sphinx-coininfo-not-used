package api

import (
	"context"

	"github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	coininfo.UnimplementedSphinxCoinInfoServer
}

func Register(server grpc.ServiceRegistrar) {
	coininfo.RegisterSphinxCoinInfoServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return coininfo.RegisterSphinxCoinInfoHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
