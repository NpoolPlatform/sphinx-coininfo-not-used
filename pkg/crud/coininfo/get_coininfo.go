package coininfo

import (
	"context"

	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/coininfo"
	constant "github.com/NpoolPlatform/sphinx-coininfo/pkg/message/const"
	"github.com/google/uuid"
)

func GetCoinInfoByID(ctx context.Context, id string) (coinInfo *ent.CoinInfo, err error) {
	coinInfo, err = db.Client().CoinInfo.Query().
		Where(coininfo.ID(uuid.MustParse(id))).
		Only(ctx)
	return
}

func GetCoinInfoByName(ctx context.Context, coinName string) (coinInfo *ent.CoinInfo, err error) {
	coinInfo, err = db.Client().CoinInfo.Query().
		Where(coininfo.Name(coinName)).
		Only(ctx)
	return
}

type GetAllCoinInfosParams struct {
	preSale       bool
	name          string
	offset, limit int
}

func GetAllCoinInfos(ctx context.Context, params GetAllCoinInfosParams) ([]*ent.CoinInfo, int, error) {
	if params.limit == 0 {
		params.limit = constant.PageSize
	}

	stm := db.Client().
		CoinInfo.
		Query()

	if params.name != "" {
		stm.Where(coininfo.NameEQ(params.name))
	}

	if params.preSale {
		stm.Where(coininfo.PreSaleEQ(params.preSale))
	}

	// total
	total, err := stm.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// infos
	coinInfos, err := stm.
		Order(ent.Desc(coininfo.FieldID)).
		Offset(params.offset).
		Limit(params.limit).
		All(ctx)

	return coinInfos, total, err
}
