package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/crud/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetCoinInfo get coininfo by id or name, id high prio
func (s *Server) GetCoinInfo(ctx context.Context, in *npool.GetCoinInfoRequest) (*npool.GetCoinInfoResponse, error) {
	if in.GetID() == "" && in.GetName() == "" {
		logger.Sugar().Errorf("GetCoinInfo check ID or Name is empty")
		return nil, status.Error(codes.InvalidArgument, "not allow ID or Name both empty")
	}

	var (
		coinInfo *ent.CoinInfo
		err      error
	)

	if in.GetID() != "" {
		if _, err := uuid.Parse(in.GetID()); err != nil {
			logger.Sugar().Errorf("GetCoinInfo check ID not a valid uuid")
			return nil, status.Error(codes.InvalidArgument, "ID invalid")
		}

		coinInfo, err = coininfo.GetCoinInfoByID(ctx, in.GetID())
		if err != nil {
			logger.Sugar().Errorf("GetCoinInfo call GetCoinInfoByID error %v", err)
			return nil, status.Error(codes.Internal, "internal server error")
		}
	} else if in.GetName() != "" {
		coinInfo, err = coininfo.GetCoinInfoByName(ctx, in.GetName())
		if err != nil {
			logger.Sugar().Errorf("GetCoinInfo call GetCoinInfoByName error %v", err)
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &npool.GetCoinInfoResponse{
		Info: &npool.CoinInfo{
			ID:      coinInfo.ID.String(),
			PreSale: coinInfo.PreSale,
			Name:    coinInfo.Name,
			Unit:    coinInfo.Unit,
			Logo:    coinInfo.Logo,
		},
	}, nil
}
