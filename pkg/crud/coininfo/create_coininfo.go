package coininfo

import (
	"context"

	"github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	"github.com/google/uuid"
)

func CreateCoinInfo(ctx context.Context, info *coininfo.CoinInfo) (*ent.CoinInfo, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return db.Client().CoinInfo.Create().
		SetID(uid).
		SetName(info.GetName()).
		SetUnit(info.GetUnit()).
		SetLogo(info.GetLogo()).
		SetPreSale(info.GetPreSale()).
		Save(ctx)
}
