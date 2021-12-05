package coininfo

import (
	"context"

	"github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db"
	dcoin "github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/coininfo"
)

func CreateCoinInfo(ctx context.Context, info *coininfo.CoinInfo) error {
	return db.Client().CoinInfo.Create().
		SetName(info.GetName()).
		SetUnit(info.GetUnit()).
		SetLogo(info.GetLogo()).
		SetPreSale(info.GetPreSale()).
		OnConflictColumns(dcoin.FieldName).
		Exec(ctx)
}
