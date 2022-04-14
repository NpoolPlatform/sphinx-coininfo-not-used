package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/crud/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/crud/description"
	ccoin "github.com/NpoolPlatform/sphinx-coininfo/pkg/message/const"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateCoinDescription ..
func (s *Server) CreateCoinDescription(ctx context.Context, in *npool.CreateCoinDescriptionRequest) (*npool.CreateCoinDescriptionResponse, error) {
	if in.GetTitle() == "" {
		logger.Sugar().Error("CreateCoinDescription check Title is empty")
		return nil, status.Error(codes.InvalidArgument, "Title empty")
	}

	if in.GetMessage() == "" {
		logger.Sugar().Error("CreateCoinDescription check Message is empty")
		return nil, status.Error(codes.InvalidArgument, "Message empty")
	}

	if in.GetUsedFor() == "" {
		logger.Sugar().Error("CreateCoinDescription check UseFor is empty")
		return nil, status.Error(codes.InvalidArgument, "UseFor empty")
	}

	if in.GetCoinTypeID() == "" {
		logger.Sugar().Error("CreateCoinDescription check CoinTypeID is empty")
		return nil, status.Error(codes.InvalidArgument, "CoinTypeID empty")
	}

	coinID, err := uuid.Parse(in.GetCoinTypeID())
	if err != nil {
		logger.Sugar().Errorf("CreateCoinDescription parse CoinTypeID: %s invalid", in.GetCoinTypeID())
		return nil, status.Error(codes.InvalidArgument, "CoinTypeID invalid")
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	existCoin, err := coininfo.ExistCoinInfoByID(ctx, coinID)
	if err != nil {
		logger.Sugar().Errorf("CreateCoinDescription call GetCoinInfoByID error %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}
	if !existCoin {
		logger.Sugar().Errorf("CreateCoinDescription call GetCoinInfoByID CoinTypeID: %v error %v", in.GetCoinTypeID(), err)
		return nil, status.Errorf(codes.NotFound, "CoinTypeID: %v not found", in.GetCoinTypeID())
	}

	existCoinDesc, err := description.ExistCoinDescriptionByCoinID(ctx, coinID)
	if err != nil {
		logger.Sugar().Errorf("CreateCoinDescription call ExistCoinDescriptionByCoinID error %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	if existCoinDesc {
		logger.Sugar().Errorf("CreateCoinDescription for CoinTypeID: %s already exist", in.GetCoinTypeID())
		return nil, status.Errorf(codes.AlreadyExists, "CoinTypeID: %v already exist", in.GetCoinTypeID())
	}

	coinDesc, err := description.CreateCoinDescription(ctx, &description.CreateCoinDescriptionParam{
		CoinTypeID: coinID,
		Title:      in.GetTitle(),
		Message:    in.GetMessage(),
		UsedFor:    in.GetUsedFor(),
	})
	if err != nil {
		logger.Sugar().Errorf("CreateCoinDescription call CreateCoinDescription error %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &npool.CreateCoinDescriptionResponse{
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
