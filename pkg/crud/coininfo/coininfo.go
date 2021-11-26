package coininfo

import (
	"context"

	//nolint
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/coininfo"
	"github.com/google/uuid"
)

func GetCoinInfoByName(ctx context.Context, coinName string) (coinInfo *ent.CoinInfo, err error) {
	coinInfo, err = db.Client().CoinInfo.Query().
		Where(coininfo.Name(coinName)).
		Only(ctx)
	return
}

func GetCoinInfoByID(ctx context.Context, id string) (coinInfo *ent.CoinInfo, err error) {
	coinInfo, err = db.Client().CoinInfo.Query().
		Where(coininfo.ID(uuid.MustParse(id))).
		Only(ctx)
	return
}

func GetAllCoinInfos(ctx context.Context) (coinInfos []*ent.CoinInfo, err error) {
	coinInfos, err = db.Client().CoinInfo.Query().All(ctx)
	return
}

func UpsertCoinInfoByName(ctx context.Context, coinInfo *ent.CoinInfo, coinName string) (coinInfoLatest *ent.CoinInfo, err error) {
	coinInfoLatest, err = db.Client().CoinInfo.Query().Where(
		coininfo.Name(coinName),
	).Only(ctx)
	if err == nil {
		if coinInfoLatest.Unit != coinInfo.Unit ||
			coinInfoLatest.IsPresale != coinInfo.IsPresale ||
			coinInfoLatest.LogoImage != coinInfo.LogoImage {
			coinInfoLatest, err = coinInfoLatest.Update().
				SetUnit(coinInfo.Unit).
				SetIsPresale(coinInfo.IsPresale).
				SetLogoImage(coinInfo.LogoImage).
				Save(ctx)
		}
	} else {
		coinInfoLatest, err = db.Client().CoinInfo.Create().
			SetID(uuid.New()).
			SetName(coinInfo.Name).
			SetUnit(coinInfo.Unit).
			SetIsPresale(coinInfo.IsPresale).
			SetLogoImage(coinInfo.LogoImage).
			SetCoinTypeID(coinInfo.CoinTypeID).
			Save(ctx)
	}
	return coinInfoLatest, err
}

func UpdatePreSaleByID(ctx context.Context, preSale bool, id string) (coinInfoNew *ent.CoinInfo, err error) {
	uid := uuid.MustParse(id)
	coinInfoNew, err = db.Client().CoinInfo.Query().Where(
		coininfo.ID(uid),
	).Only(ctx)
	if err == nil {
		coinInfoNew, err = coinInfoNew.Update().SetIsPresale(preSale).Save(ctx)
	}
	return
}
