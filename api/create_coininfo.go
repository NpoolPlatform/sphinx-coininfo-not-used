package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/price"
	npool "github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/crud/coininfo"
	ccoin "github.com/NpoolPlatform/sphinx-coininfo/pkg/message/const"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateCoinInfo support duplicate
func (s *Server) CreateCoinInfo(ctx context.Context, in *npool.CreateCoinInfoRequest) (*npool.CreateCoinInfoResponse, error) {
	_, span := otel.Tracer(ccoin.ServiceName).Start(ctx, "CreateCoinInfo")
	defer span.End()

	span.SetAttributes(
		attribute.Bool("PreSale", in.GetPreSale()),
		attribute.String("Name", in.GetName()),
		attribute.String("Unit", in.GetUnit()),
		attribute.String("Logo", in.GetLogo()),
		attribute.String("ENV", in.GetENV()),
	)

	if in.GetName() == "" {
		logger.Sugar().Error("CreateCoinInfo check Name is empty")
		return nil, status.Error(codes.InvalidArgument, "Name empty")
	}

	if in.GetUnit() == "" {
		logger.Sugar().Error("CreateCoinInfo check Unit is empty")
		return nil, status.Error(codes.InvalidArgument, "Unit empty")
	}

	if in.GetENV() == "" {
		logger.Sugar().Error("CreateCoinInfo check ENV is empty")
		return nil, status.Error(codes.InvalidArgument, "ENV empty")
	}

	if in.GetENV() != "main" && in.GetENV() != "test" {
		logger.Sugar().Errorf("CreateCoinInfo check ENV: %s invalid", in.GetENV())
		return nil, status.Error(codes.InvalidArgument, "ENV invalid")
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	span.AddEvent("call db CreateCoinInfo",
		trace.WithAttributes(
			attribute.Bool("PreSale", in.GetPreSale()),
			attribute.String("Name", in.GetName()),
			attribute.String("Unit", in.GetUnit()),
			attribute.String("Logo", in.GetLogo()),
			attribute.String("ENV", in.GetENV()),
		),
	)
	err := coininfo.CreateCoinInfo(ctx, &npool.CoinInfo{
		PreSale: in.GetPreSale(),
		Name:    in.GetName(),
		Unit:    in.GetUnit(),
		Logo:    in.GetLogo(),
		ENV:     in.GetENV(),
	})
	if err != nil {
		logger.Sugar().Errorf("CreateCoinInfo call CreateCoinInfo error %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	span.AddEvent("call db GetCoinInfoByName", trace.WithAttributes(
		attribute.String("Name", in.GetName()),
	))
	coinInfo, err := coininfo.GetCoinInfoByName(ctx, in.GetName())
	if err != nil {
		logger.Sugar().Errorf("CreateCoinInfo call GetCoinInfoByName error %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &npool.CreateCoinInfoResponse{
		Info: &npool.CoinInfo{
			ID:             coinInfo.ID.String(),
			PreSale:        coinInfo.PreSale,
			ForPay:         coinInfo.ForPay,
			Name:           coinInfo.Name,
			Unit:           coinInfo.Unit,
			Logo:           coinInfo.Logo,
			ENV:            coinInfo.Env,
			ReservedAmount: price.DBPriceToVisualPrice(coinInfo.ReservedAmount),
			CreatedAt:      coinInfo.CreatedAt,
			UpdatedAt:      coinInfo.UpdatedAt,
		},
	}, nil
}
