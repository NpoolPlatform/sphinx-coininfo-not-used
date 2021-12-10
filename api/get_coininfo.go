package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/price"
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
		err      error
		coinInfo *ent.CoinInfo
	)

	if in.GetID() != "" {
		id, err := uuid.Parse(in.GetID())
		if err != nil {
			logger.Sugar().Errorf("GetCoinInfo check ID: %v not a valid uuid", in.GetID())
			return nil, status.Errorf(codes.InvalidArgument, "ID: %v invalid", in.GetID())
		}

		coinInfo, err = coininfo.GetCoinInfoByID(ctx, id)
		if ent.IsNotFound(err) {
			logger.Sugar().Errorf("GetCoinInfo call GetCoinInfoByID ID: %v not found", in.GetID())
			return nil, status.Errorf(codes.NotFound, "ID: %v not found", in.GetID())
		}

		if err != nil {
			logger.Sugar().Errorf("GetCoinInfo call GetCoinInfoByID error %v", err)
			return nil, status.Error(codes.Internal, "internal server error")
		}
	} else if in.GetName() != "" {
		coinInfo, err = coininfo.GetCoinInfoByName(ctx, in.GetName())
		if ent.IsNotFound(err) {
			logger.Sugar().Errorf("GetCoinInfo call GetCoinInfoByName Name: %v not found", in.GetName())
			return nil, status.Errorf(codes.NotFound, "Name: %v not found", in.GetName())
		}
		if err != nil {
			logger.Sugar().Errorf("GetCoinInfo call GetCoinInfoByName error %v", err)
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &npool.GetCoinInfoResponse{
		Info: &npool.CoinInfo{
			ID:             coinInfo.ID.String(),
			PreSale:        coinInfo.PreSale,
			Name:           coinInfo.Name,
			Unit:           coinInfo.Unit,
			ReservedAmount: price.DBPriceToVisualPrice(coinInfo.ReservedAmount),
			Logo:           coinInfo.Logo,
			CreatedAt:      coinInfo.CreatedAt,
			UpdatedAt:      coinInfo.UpdatedAt,
		},
	}, nil
}
