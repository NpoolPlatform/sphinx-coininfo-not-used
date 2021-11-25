package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/coininfo" //nolint
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/app"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	FlagOutputError bool
	errDefault      error
)

func init() {
	FlagOutputError = true
	errDefault = status.Error(codes.Internal, "internal server error")
}

func (s *Server) GetCoinInfos(ctx context.Context, in *emptypb.Empty) (resp *npool.GetCoinInfosResponse, err error) {
	resp, err = app.GetCoinInfos(ctx, &npool.GetCoinInfosRequest{})
	if err != nil {
		err = LogWhenError(err, "getcoininfos error")
	}
	return
}

func (s *Server) GetCoinInfo(ctx context.Context, in *npool.GetCoinInfoRequest) (resp *npool.GetCoinInfoResponse, err error) {
	resp, err = app.GetCoinInfo(ctx, in)
	if err != nil {
		err = LogWhenError(err, "getcoininfo error")
	}
	return
}

func (s *Server) CreateCoinInfo(ctx context.Context, in *npool.CreateCoinInfoRequest) (resp *npool.CreateCoinInfoResponse, err error) {
	resp, err = app.CreateCoinInfo(ctx, in)
	if err != nil {
		err = LogWhenError(err, "createcoininfo error")
	}
	return
}

func (s *Server) UpdateCoinInfo(ctx context.Context, in *npool.UpdateCoinInfoRequest) (resp *npool.UpdateCoinInfoResponse, err error) {
	resp, err = app.UpdateCoinInfo(ctx, in)
	if err != nil {
		err = LogWhenError(err, "updatecoininfo error")
	}
	return
}

func LogWhenError(err error, msg string) (errNew error) {
	if err != nil {
		logger.Sugar().Errorf(msg+" grpc error: %v", err)
		if FlagOutputError {
			errNew = errDefault
		} else {
			errNew = status.Errorf(codes.Internal, msg+" grpc error: %v", err)
		}
	}
	return
}
