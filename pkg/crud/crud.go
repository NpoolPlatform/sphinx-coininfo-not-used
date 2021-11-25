package crud

import (
	"context"

	//nolint
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/coininfo"
	"github.com/google/uuid"
)

func GetInfo(ctx context.Context, id string) (coinInfo *ent.CoinInfo, err error) {
	tmpID := uuid.MustParse(id)
	coinInfo, err = db.Client().CoinInfo.Query().
		Where(coininfo.ID(tmpID)).
		Only(ctx)
	return
}

func GetInfos(ctx context.Context) (coinInfos []*ent.CoinInfo, err error) {
	coinInfos, err = db.Client().CoinInfo.Query().All(ctx)
	return
}

// 当前逻辑： INSERT OR UPDATE
func CreateCoin(ctx context.Context, coinInfo *ent.CoinInfo) (coinInfoCreated *ent.CoinInfo, err error) {
	coinInfoCreated, err = db.Client().CoinInfo.Query().Where(
		coininfo.CoinTypeID(coinInfo.CoinTypeID),
	).Only(ctx)
	if err == nil {
		coinInfoCreated, err = coinInfoCreated.Update().
			SetName(coinInfo.Name).
			SetUnit(coinInfo.Unit).
			SetIsPresale(coinInfo.IsPresale).
			SetLogoImage(coinInfo.LogoImage).
			Save(ctx)
	} else {
		coinInfoCreated, err = db.Client().CoinInfo.Create().
			SetID(uuid.New()).
			SetName(coinInfo.Name).
			SetUnit(coinInfo.Unit).
			SetIsPresale(coinInfo.IsPresale).
			SetLogoImage(coinInfo.LogoImage).
			Save(ctx)
	}
	return coinInfoCreated, err
}

func UpdatePreSale(ctx context.Context, coinInfo *ent.CoinInfo) (coinInfoNew *ent.CoinInfo, err error) {
	coinInfoNew, err = db.Client().CoinInfo.Query().Where(
		coininfo.ID(coinInfo.ID),
	).Only(ctx)
	if err == nil {
		coinInfoNew, err = coinInfoNew.Update().SetIsPresale(coinInfo.IsPresale).Save(ctx)
	}
	return
}
