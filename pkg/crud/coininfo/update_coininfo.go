package coininfo

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	"github.com/google/uuid"
)

func UpdateCoinInfoByID(ctx context.Context, preSale bool, logo, id string, reservedAmount float64) (coinInfo *ent.CoinInfo, err error) {
	stmt := db.Client().
		CoinInfo.
		UpdateOneID(uuid.MustParse(id))

	if preSale {
		stmt.SetPreSale(preSale)
	}
	if logo != "" {
		stmt.SetLogo(logo)
	}
	if reservedAmount > 0 {
		stmt.SetReservedAmount(price.VisualPriceToDBPrice(reservedAmount))
	}
	coinInfo, err = stmt.
		Save(ctx)
	return
}
