package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/crud/coininfo"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateCoinInfo(ctx context.Context, in *npool.UpdateCoinInfoRequest) (*npool.UpdateCoinInfoResponse, error) {
	if in.GetID() == "" {
		logger.Sugar().Errorf("UpdateCoinInfo check ID is empty")
		return nil, status.Error(codes.InvalidArgument, "ID empty")
	}

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorf("UpdateCoinInfo check ID not a valid uuid")
		return nil, status.Error(codes.InvalidArgument, "ID invalid")
	}

	_, err := coininfo.UpdateCoinInfoByID(ctx, in.GetPreSale(), in.GetLogo(), in.GetID())
	if err != nil {
		logger.Sugar().Errorf("UpdateCoinInfo call UpdateCoinInfoByID error %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &npool.UpdateCoinInfoResponse{}, nil
}
