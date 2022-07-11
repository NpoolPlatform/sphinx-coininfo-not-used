package coininfo

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	ccoin "github.com/NpoolPlatform/sphinx-coininfo/pkg/message/const"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

func UpdateCoinInfoByID(ctx context.Context, preSale, forPay bool, logo, id, homePage, specs string, reservedAmount float64) (*ent.CoinInfo, error) {
	_, span := otel.Tracer(ccoin.ServiceName).Start(ctx, "UpdateCoinInfoByID")
	defer span.End()

	span.SetAttributes(
		attribute.Bool("PreSale", preSale),
		attribute.Bool("ForPay", forPay),
		attribute.String("Logo", logo),
		attribute.String("ID", id),
		attribute.String("HomePage", homePage),
		attribute.String("Specs", specs),
		attribute.Float64("ReservedAmount", reservedAmount),
	)

	client, err := db.Client()
	if err != nil {
		span.SetStatus(codes.Error, "get db client fail")
		span.RecordError(err)
		return nil, err
	}

	stmt := client.
		CoinInfo.
		UpdateOneID(uuid.MustParse(id))

	stmt.SetPreSale(preSale)
	stmt.SetForPay(forPay)
	if logo != "" {
		stmt.SetLogo(logo)
	}
	if homePage != "" {
		stmt.SetHomePage(homePage)
	}
	if specs != "" {
		stmt.SetSpecs(specs)
	}
	if reservedAmount > 0 {
		stmt.SetReservedAmount(price.VisualPriceToDBPrice(reservedAmount))
	}

	return stmt.
		Save(ctx)
}
