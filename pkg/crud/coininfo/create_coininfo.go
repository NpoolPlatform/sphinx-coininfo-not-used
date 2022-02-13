package coininfo

import (
	"context"

	"github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db"
	dcoin "github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/coininfo"
)

func CreateCoinInfo(ctx context.Context, info *coininfo.CoinInfo) error {
	client, err := db.Client()
	if err != nil {
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
