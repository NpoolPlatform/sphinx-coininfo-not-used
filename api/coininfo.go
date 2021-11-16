package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/crud"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (Server) GetCoinInfos(ctx context.Context, in *npool.GetCoinInfosRequest) (resp *npool.GetCoinInfosResponse, err error) {
	return crud.GetCoinInfos(ctx, in)
}

func (Server) RegisterCoin(ctx context.Context, in *npool.RegisterCoinRequest) (resp *npool.RegisterCoinResponse, err error) {
	resp, err = crud.RegisterCoin(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("register coin error: %w", err)
		return &npool.RegisterCoinResponse{Info: "failed"}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
