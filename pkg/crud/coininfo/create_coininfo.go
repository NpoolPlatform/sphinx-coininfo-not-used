package coininfo

import (
	"context"

	"github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db"
	dcoin "github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/coininfo"
	"github.com/google/uuid"
)

func CreateCoinInfo(ctx context.Context, info *coininfo.CoinInfo) (uuid.UUID, error) {
	return db.Client().CoinInfo.Create().
		SetName(info.GetName()).
		SetUnit(info.GetUnit()).
		SetLogo(info.GetLogo()).
		SetPreSale(info.GetPreSale()).
		OnConflictColumns(dcoin.FieldName).
		ID(ctx)
}
