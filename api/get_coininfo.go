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

func (s *Server) GetCoinInfo(ctx context.Context, in *npool.GetCoinInfoRequest) (*npool.GetCoinInfoResponse, error) {
	if in.GetID() == "" {
		logger.Sugar().Errorf("GetCoinInfo check ID is empty")
		return nil, status.Error(codes.InvalidArgument, "ID empty")
	}

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorf("GetCoinInfo check ID not a valid uuid")
		return nil, status.Error(codes.InvalidArgument, "ID invalid")
	}

	resp, err := coininfo.GetCoinInfoByID(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorf("GetCoinInfo call GetCoinInfoByID error %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &npool.GetCoinInfoResponse{
		Info: &npool.CoinInfo{
			ID:      resp.ID.String(),
			PreSale: resp.PreSale,
			Name:    resp.Name,
			Unit:    resp.Unit,
			Logo:    resp.Logo,
		},
	}, nil
}
