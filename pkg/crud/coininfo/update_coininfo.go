package coininfo

import (
	"context"

	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	"github.com/google/uuid"
)

func UpdateCoinInfoByID(ctx context.Context, preSale bool, logo string, id string) (coinInfo *ent.CoinInfo, err error) {
	stmt := db.Client().
		CoinInfo.
		UpdateOneID(uuid.MustParse(id))

	if preSale {
		stmt.SetPreSale(preSale)
	}
	if logo != "" {
		stmt.SetLogo(logo)
	}

	coinInfo, err = stmt.
		Save(ctx)
	return
}
