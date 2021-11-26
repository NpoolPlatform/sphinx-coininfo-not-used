package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/coininfo" //nolint
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/middleware"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	FlagOutputError = true
	errDefault      = status.Error(codes.Internal, "internal server error")
)

func (s *Server) GetCoinInfo(ctx context.Context, in *npool.GetCoinInfoRequest) (*npool.GetCoinInfoResponse, error) {
	resp, err := middleware.GetCoinInfo(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("getcoininfo error %v \n when %+v", err, in)
		return &npool.GetCoinInfoResponse{}, patchError(err)
	}
	return resp, nil
}

func (s *Server) GetCoinInfos(ctx context.Context, in *emptypb.Empty) (*npool.GetCoinInfosResponse, error) {
	resp, err := middleware.GetCoinInfos(ctx, &npool.GetCoinInfosRequest{})
	if err != nil {
		logger.Sugar().Errorf("getcoininfos error %v \n when %+v", err, in)
		return &npool.GetCoinInfosResponse{}, patchError(err)
	}
	return resp, nil
}

func (s *Server) CreateCoinInfo(ctx context.Context, in *npool.CreateCoinInfoRequest) (*npool.CreateCoinInfoResponse, error) {
	resp, err := middleware.CreateCoinInfo(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("createcoininfo error %v \n when %+v", err, in)
		return &npool.CreateCoinInfoResponse{}, patchError(err)
	}
	return resp, nil
}

func (s *Server) UpdateCoinInfo(ctx context.Context, in *npool.UpdateCoinInfoRequest) (*npool.UpdateCoinInfoResponse, error) {
	resp, err := middleware.UpdateCoinInfo(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("updatecoininfo error %v \n when %+v", err, in)
		return &npool.UpdateCoinInfoResponse{}, patchError(err)
	}
	return resp, nil
}

func patchError(err error) (errGRPC error) {
	if !FlagOutputError {
		errGRPC = errDefault
	} else if _, ok := status.FromError(err); ok {
		errGRPC = err
	} else {
		errGRPC = status.Error(codes.Internal, err.Error())
	}
	return
}
