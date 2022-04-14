package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/price"
	npool "github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/crud/coininfo"
	ccoin "github.com/NpoolPlatform/sphinx-coininfo/pkg/message/const"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateCoinInfo(ctx context.Context, in *npool.UpdateCoinInfoRequest) (*npool.UpdateCoinInfoResponse, error) {
	if in.GetID() == "" {
		logger.Sugar().Errorf("UpdateCoinInfo check ID is empty")
		return nil, status.Error(codes.InvalidArgument, "ID empty")
	}

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorf("UpdateCoinInfo check ID not a valid uuid")
		return nil, status.Error(codes.InvalidArgument, "ID invalid")
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	existCoin, err := coininfo.ExistCoinInfoByID(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("UpdateCoinInfo call GetCoinInfoByID error %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	if !existCoin {
		logger.Sugar().Errorf("UpdateCoinInfo call GetCoinInfoByID ID: %v not found", in.GetID())
		return nil, status.Errorf(codes.NotFound, "ID: %v not found", in.GetID())
	}

	coinInfo, err := coininfo.UpdateCoinInfoByID(ctx, in.GetPreSale(), in.GetForPay(), in.GetLogo(), in.GetID(), in.GetHomePage(), in.GetSpecs(), in.GetReservedAmount())
	if err != nil {
		logger.Sugar().Errorf("UpdateCoinInfo call UpdateCoinInfoByID error %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &npool.UpdateCoinInfoResponse{
		Info: &npool.CoinInfo{
			ID:             coinInfo.ID.String(),
			PreSale:        coinInfo.PreSale,
			ForPay:         coinInfo.ForPay,
			Name:           coinInfo.Name,
			Unit:           coinInfo.Unit,
			Logo:           coinInfo.Logo,
			HomePage:       coinInfo.HomePage,
			Specs:          coinInfo.Specs,
			ReservedAmount: price.DBPriceToVisualPrice(coinInfo.ReservedAmount),
			CreatedAt:      coinInfo.CreatedAt,
			UpdatedAt:      coinInfo.UpdatedAt,
		},
	}, nil
}
