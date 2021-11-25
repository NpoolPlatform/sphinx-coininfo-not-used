package api

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/coininfo" //nolint
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedSphinxCoinInfoServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterSphinxCoinInfoServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterSphinxCoinInfoHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
