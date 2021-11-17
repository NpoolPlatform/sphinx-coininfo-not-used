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

func (Server) GetCoinInfo(ctx context.Context, in *npool.GetCoinInfoRequest) (resp *npool.CoinInfoRow, err error) {
	resp, err = crud.GetCoinInfo(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get coininfo error: %w", err)
		if DebugFlag {
			err = status.Error(codes.Internal, "internal server error")
		}
	}
	return
}

func (Server) RegisterCoin(ctx context.Context, in *npool.RegisterCoinRequest) (resp *npool.CoinInfoRow, err error) {
	resp, err = crud.RegisterCoin(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("register coin error: %w", err)
		if DebugFlag {
			err = status.Error(codes.Internal, "internal server error")
		}
	}
	return
}

func (Server) SetCoinPresale(ctx context.Context, in *npool.SetCoinPresaleRequest) (resp *npool.CoinInfoRow, err error) {
	resp, err = crud.SetCoinPresale(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("set coinpresale error: %w", err)
		if DebugFlag {
			err = status.Error(codes.Internal, "internal server error")
		}
	}
	return
}
