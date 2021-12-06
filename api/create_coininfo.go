package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/crud/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCoinInfo(ctx context.Context, in *npool.CreateCoinInfoRequest) (*npool.CreateCoinInfoResponse, error) {
	if in.GetName() == "" {
		logger.Sugar().Errorf("CreateCoinInfo check Name is empty")
		return nil, status.Error(codes.InvalidArgument, "Name empty")
	}

	_, err := coininfo.GetCoinInfoByName(ctx, in.GetName())
	if !ent.IsNotFound(err) {
		logger.Sugar().Errorf("CreateCoinInfo check Name: %v already exist", in.GetName())
		return nil, status.Errorf(codes.AlreadyExists, "Name: %v already exist", in.GetName())
	}

	if err != nil {
		logger.Sugar().Errorf("CreateCoinInfo call GetCoinInfoByName error %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	if in.GetUnit() == "" {
		logger.Sugar().Errorf("CreateCoinInfo check Unit is empty")
		return nil, status.Error(codes.InvalidArgument, "Unit empty")
	}

	err = coininfo.CreateCoinInfo(ctx, &npool.CoinInfo{
		PreSale: in.GetPreSale(),
		Name:    in.GetName(),
		Unit:    in.GetUnit(),
		Logo:    in.GetLogo(),
	})
	if err != nil {
		logger.Sugar().Errorf("CreateCoinInfo call CreateCoinInfo error %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &npool.CreateCoinInfoResponse{}, nil
}
