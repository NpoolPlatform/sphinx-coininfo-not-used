package coininfo

import (
	"context"

	"github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db"
	dcoin "github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/coininfo"
	ccoin "github.com/NpoolPlatform/sphinx-coininfo/pkg/message/const"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

func CreateCoinInfo(ctx context.Context, info *coininfo.CoinInfo) error {
	_, span := otel.Tracer(ccoin.ServiceName).Start(ctx, "CreateCoinInfo")
	defer span.End()

	span.SetAttributes(
		attribute.Bool("PreSale", info.PreSale),
		attribute.String("Name", info.Name),
		attribute.String("Unit", info.Unit),
		attribute.String("Logo", info.Logo),
		attribute.String("ENV", info.ENV),
	)

	client, err := db.Client()
	if err != nil {
		span.SetStatus(codes.Error, "get db client fail")
		span.RecordError(err)
		return err
	}

	return client.CoinInfo.Create().
		SetName(info.GetName()).
		SetUnit(info.GetUnit()).
		SetLogo(info.GetLogo()).
		SetEnv(info.GetENV()).
		SetPreSale(info.GetPreSale()).
		SetForPay(false).
		OnConflictColumns(dcoin.FieldName).
		Ignore().
		Exec(ctx)
}
