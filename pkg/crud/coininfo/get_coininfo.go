package coininfo

import (
	"context"

	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/coininfo"
	ccoin "github.com/NpoolPlatform/sphinx-coininfo/pkg/message/const"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

func GetCoinInfoByID(ctx context.Context, id uuid.UUID) (coinInfo *ent.CoinInfo, err error) {
	_, span := otel.Tracer(ccoin.ServiceName).Start(ctx, "GetCoinInfoByID")
	defer span.End()

	span.SetAttributes(
		attribute.String("ID", id.String()),
	)

	client, err := db.Client()
	if err != nil {
		span.SetStatus(codes.Error, "get db client fail")
		span.RecordError(err)
		return nil, err
	}

	coinInfo, err = client.CoinInfo.Query().
		Where(coininfo.ID(id)).
		Only(ctx)
	return
}

func GetCoinInfoByName(ctx context.Context, coinName string) (coinInfo *ent.CoinInfo, err error) {
	_, span := otel.Tracer(ccoin.ServiceName).Start(ctx, "GetCoinInfoByName")
	defer span.End()

	span.SetAttributes(
		attribute.String("CoinName", coinName),
	)

	client, err := db.Client()
	if err != nil {
		span.SetStatus(codes.Error, "get db client fail")
		span.RecordError(err)
		return nil, err
	}

	coinInfo, err = client.CoinInfo.Query().
		Where(coininfo.Name(coinName)).
		Only(ctx)
	return
}

func ExistCoinInfoByID(ctx context.Context, coinID uuid.UUID) (bool, error) {
	_, span := otel.Tracer(ccoin.ServiceName).Start(ctx, "ExistCoinInfoByID")
	defer span.End()

	span.SetAttributes(
		attribute.String("CoinID", coinID.String()),
	)

	client, err := db.Client()
	if err != nil {
		span.SetStatus(codes.Error, "get db client fail")
		span.RecordError(err)
		return false, err
	}

	return client.CoinInfo.Query().
		Where(coininfo.IDEQ(coinID)).
		Exist(ctx)
}

type GetAllCoinInfosParams struct {
	PreSale       bool
	Name          string
	Offset, Limit int
}

func GetAllCoinInfos(ctx context.Context, params GetAllCoinInfosParams) ([]*ent.CoinInfo, int, error) {
	_, span := otel.Tracer(ccoin.ServiceName).Start(ctx, "GetAllCoinInfos")
	defer span.End()

	span.SetAttributes(
		attribute.Bool("PreSale", params.PreSale),
		attribute.String("Name", params.Name),
		attribute.Int("Offset", params.Offset),
		attribute.Int("Limit", params.Limit),
	)

	if params.Limit == 0 {
		params.Limit = ccoin.PageSize
	}

	client, err := db.Client()
	if err != nil {
		span.SetStatus(codes.Error, "get db client fail")
		span.RecordError(err)
		return nil, 0, err
	}

	stm := client.
		CoinInfo.
		Query()

	if params.Name != "" {
		stm.Where(coininfo.NameEQ(params.Name))
	}

	if params.PreSale {
		stm.Where(coininfo.PreSaleEQ(params.PreSale))
	}

	// total
	total, err := stm.Count(ctx)
	if err != nil {
		span.SetStatus(codes.Error, "call count fail")
		span.RecordError(err)
		return nil, 0, err
	}

	// infos
	coinInfos, err := stm.
		Order(ent.Desc(coininfo.FieldCreatedAt)).
		Offset(params.Offset).
		Limit(params.Limit).
		All(ctx)

	return coinInfos, total, err
}
