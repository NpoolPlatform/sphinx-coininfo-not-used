package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/crud"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (Server) GetCoinInfos(ctx context.Context, in *emptypb.Empty) (resp *npool.GetCoinInfosResponse, err error) {
	return crud.GetCoinInfos(ctx, &npool.GetCoinInfosRequest{})
}

func (Server) GetCoinInfo(ctx context.Context, in *npool.GetCoinInfoRequest) (resp *npool.GetCoinInfoResponse, err error) {
	resp, err = crud.GetCoinInfo(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get coininfo error: %w", err)
		if DebugFlag {
			err = status.Error(codes.Internal, "internal server error")
		}
	}
	return
}

func (Server) CreateCoinInfo(ctx context.Context, in *npool.CreateCoinInfoRequest) (resp *npool.CreateCoinInfoResponse, err error) {
	resp, err = crud.RegisterCoin(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("register coin error: %w", err)
		if DebugFlag {
			err = status.Error(codes.Internal, "internal server error")
		}
	}
	return
}

func (Server) UpdateCoinInfo(ctx context.Context, in *npool.UpdateCoinInfoRequest) (resp *npool.UpdateCoinInfoResponse, err error) {
	resp, err = crud.SetCoinPresale(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("set coinpresale error: %w", err)
		if DebugFlag {
			err = status.Error(codes.Internal, "internal server error")
		}
	}
	return
}
