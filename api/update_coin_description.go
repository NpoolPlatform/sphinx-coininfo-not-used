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

func (s *Server) UpdateCoinDescription(ctx context.Context, in *npool.UpdateCoinDescriptionRequest) (*npool.UpdateCoinDescriptionResponse, error) {
	_, span := otel.Tracer(ccoin.ServiceName).Start(ctx, "CreateCoinInfo")
	defer span.End()

	if in.GetTitle() == "" {
		logger.Sugar().Error("UpdateCoinDescription check Title is empty")
		return nil, status.Error(codes.InvalidArgument, "Title empty")
	}

	if in.GetMessage() == "" {
		logger.Sugar().Error("UpdateCoinDescription check Message is empty")
		return nil, status.Error(codes.InvalidArgument, "Message empty")
	}

	if in.GetUsedFor() == "" {
		logger.Sugar().Error("UpdateCoinDescription check UseFor is empty")
		return nil, status.Error(codes.InvalidArgument, "UseFor empty")
	}

	if in.GetCoinTypeID() == "" {
		logger.Sugar().Error("UpdateCoinDescription check CoinTypeID is empty")
		return nil, status.Error(codes.InvalidArgument, "CoinTypeID empty")
	}

	coinID, err := uuid.Parse(in.GetCoinTypeID())
	if err != nil {
		logger.Sugar().Errorf("UpdateCoinDescription parse CoinTypeID: %s invalid", in.GetCoinTypeID())
		return nil, status.Error(codes.InvalidArgument, "CoinTypeID invalid")
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	oldCoinDesc, err := description.GetCoinDescriptionByCoinID(ctx, coinID)
	if err != nil {
		logger.Sugar().Errorf("UpdateCoinDescription call ExistCoinDescriptionByCoinID error %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	if ent.IsNotFound(err) {
		logger.Sugar().Errorf("UpdateCoinDescription call ExistCoinDescriptionByCoinID ID: %v not found", in.GetCoinTypeID())
		return nil, status.Errorf(codes.NotFound, "ID: %v not found", in.GetCoinTypeID())
	}

	coinDesc, err := description.UpdateCoinDescriptionByID(ctx, in.GetTitle(), in.GetMessage(), in.GetUsedFor(), oldCoinDesc.ID)
	if err != nil {
		logger.Sugar().Errorf("UpdateCoinDescription call UpdateCoinDescriptionByID error %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &npool.UpdateCoinDescriptionResponse{
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
