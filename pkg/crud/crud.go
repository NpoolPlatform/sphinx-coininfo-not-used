package crud

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/coininfo"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func dbRowToCoinInfoRow(row *ent.CoinInfo) *npool.CoinInfoRow {
	return &npool.CoinInfoRow{
		CoinType:  sphinxplugin.CoinType(row.ID),
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
	tmpID := int32(in.GetCoinType())
	entResp, err := db.Client().CoinInfo.Query().Where(
		coininfo.ID(tmpID),
	).First(ctx)
	if entResp == nil || err != nil {
		err = status.Errorf(codes.NotFound, "record not found, err: %v", err)
		resp = nil
		return
	}
	resp = dbRowToCoinInfoRow(entResp)
	return
}

// 注册币种
func RegisterCoin(ctx context.Context, in *npool.RegisterCoinRequest) (resp *npool.CoinInfoRow, err error) {
	fmt.Println(in)
	resp = nil
	tmpID := int32(in.CoinType.Number())
	entResp, err := db.Client().CoinInfo.Query().
		Where(
			coininfo.ID(tmpID),
		).First(ctx)
	if entResp != nil {
		// 记录已存在
		if in.Unit == entResp.Unit && in.Name == entResp.Name {
			resp = dbRowToCoinInfoRow(entResp)
			err = nil
		} else {
			err = status.Errorf(codes.AlreadyExists, "coin name already registered as: %v, unit: %v", entResp.Name, entResp.Unit)
		}
		return
	}
	if err != nil {
		fmt.Print(err)
	}
	// MARK 默认均为在售商品？
	entResp, err = db.Client().CoinInfo.Create().
		SetID(tmpID).
		SetName(in.Name).
		SetUnit(in.Unit).
		SetIsPresale(false).
		Save(ctx)
	if err == nil {
		resp = dbRowToCoinInfoRow(entResp)
	}
	return resp, err
}

// 设置币种权限
func SetCoinPresale(ctx context.Context, in *npool.SetCoinPresaleRequest) (resp *npool.CoinInfoRow, err error) {
	ci, err := db.Client().CoinInfo.Query().
		Where(
			coininfo.And(
				coininfo.ID(int32(in.CoinType)),
			),
		).First(ctx)
	if ci == nil || err != nil {
		err = status.Errorf(codes.NotFound, "no record found, err: %v", err)
		return
	}
	ci, err = ci.Update().
		SetIsPresale(in.IsPresale).
		Save(ctx)
	if err == nil {
		resp = dbRowToCoinInfoRow(ci)
	}
	return
}
