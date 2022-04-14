package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/crud/description"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	ccoin "github.com/NpoolPlatform/sphinx-coininfo/pkg/message/const"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetCoinDescription ..
func (s *Server) GetCoinDescription(ctx context.Context, in *npool.GetCoinDescriptionRequest) (*npool.GetCoinDescriptionResponse, error) {
	_, span := otel.Tracer(ccoin.ServiceName).Start(ctx, "CreateCoinInfo")
	defer span.End()

	if in.GetCoinTypeID() == "" {
		logger.Sugar().Errorf("GetCoinDescription check CoinTypeID is empty")
		return nil, status.Error(codes.InvalidArgument, "CoinTypeID is empty")
	}

	coinID, err := uuid.Parse(in.GetCoinTypeID())
	if err != nil {
		logger.Sugar().Errorf("GetCoinDescription parse CoinTypeID: %s invalid", in.GetCoinTypeID())
		return nil, status.Error(codes.InvalidArgument, "CoinTypeID invalid")
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	coinDesc, err := description.GetCoinDescriptionByCoinID(ctx, coinID)
	if ent.IsNotFound(err) {
		logger.Sugar().Errorf("GetCoinDescription call GetCoinDescriptionByCoinID CoinTypeID: %v not found", in.GetCoinTypeID())
		return nil, status.Errorf(codes.NotFound, "CoinTypeID: %v not found", in.GetCoinTypeID())
	}
	if err != nil {
		logger.Sugar().Errorf("GetCoinDescription call GetCoinDescriptionByCoinID error %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &npool.GetCoinDescriptionResponse{
		Info: &npool.CoinDescriptionInfo{
			ID:         coinDesc.ID.String(),
			CoinTypeID: coinDesc.CoinTypeID.String(),
			Title:      coinDesc.Title,
			Message:    coinDesc.Message,
			UsedFor:    coinDesc.UsedFor,
			CreatedAt:  coinDesc.CreatedAt,
			UpdatedAt:  coinDesc.UpdatedAt,
		},
	}, nil
}
