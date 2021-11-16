package crud

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/coininfo"
	"github.com/gogo/status"
	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
)

func dbRowToCoinInfoRow(row *ent.CoinInfo) *npool.CoinInfoRow {
	return &npool.CoinInfoRow{
		Id:        row.ID,
		Name:      row.Name,
		Unit:      row.Unit,
		IsPresale: row.IsPresale,
	}
}

// 查询全部币种
func GetCoinInfos(ctx context.Context, _ *npool.GetCoinInfosRequest) (resp *npool.GetCoinInfosResponse, err error) {
	entResp, err := db.Client().CoinInfo.Query().All(ctx)
	coininfos := make([]*npool.CoinInfoRow, len(entResp))
	for i, row := range entResp {
		coininfos[i] = dbRowToCoinInfoRow(row)
	}
	resp = &npool.GetCoinInfosResponse{
		Infos: coininfos,
	}
	return
}

// 获取单个币种
func GetCoinInfo(ctx context.Context, in *npool.GetCoinInfoRequest) (resp *npool.CoinInfoRow, err error) {
	entResp, err := db.Client().CoinInfo.Query().Where(
		coininfo.ID(in.CoinId),
	).First(ctx)
	if err != nil {
		err = status.Error(codes.NotFound, "record not found")
	}
	resp = dbRowToCoinInfoRow(entResp)
	return
}

// 注册币种
func RegisterCoin(ctx context.Context, in *npool.RegisterCoinRequest) (resp *npool.RegisterCoinResponse, err error) {
	entResp, err := db.Client().CoinInfo.Query().
		Where(
			coininfo.Name(in.Name),
		).First(ctx)
	if err != nil {
		err = xerrors.Errorf("internal server error: %v", err)
		return
	}
	if entResp != nil {
		// already have record
		if in.Unit == entResp.Unit {
			resp = &npool.RegisterCoinResponse{Info: "success"}
			err = nil
		} else {
			err = xerrors.Errorf("coin name already registered as: %v, unit: %v", entResp.Name, entResp.Unit)
		}
	} else {
		// MARK 默认均为在售商品？
		_, err = db.Client().CoinInfo.Create().
			SetName(in.Name).
			SetUnit(in.Unit).
			SetIsPresale(false).
			Save(ctx)
		if err == nil {
			resp = &npool.RegisterCoinResponse{Info: "success"}
		}
	}
	return
}
