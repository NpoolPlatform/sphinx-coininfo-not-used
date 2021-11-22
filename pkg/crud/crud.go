package crud

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/coininfo" //nolint
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/coininfo"
	"github.com/gogo/status"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
)

func dbRowToCoinInfo(row *ent.CoinInfo) *npool.CoinInfo {
	return &npool.CoinInfo{
		ID:      row.ID.String(),
		Enum:    row.CoinTypeID,
		Name:    row.Name,
		Unit:    row.Unit,
		PreSale: row.IsPresale,
	}
}

// 查询全部币种
func GetCoinInfos(ctx context.Context, _ *npool.GetCoinInfosRequest) (resp *npool.GetCoinInfosResponse, err error) {
	entResp, err := db.Client().CoinInfo.Query().All(ctx)
	coininfos := make([]*npool.CoinInfo, len(entResp))
	for i, row := range entResp {
		coininfos[i] = dbRowToCoinInfo(row)
	}
	resp = &npool.GetCoinInfosResponse{
		Infos: coininfos,
	}
	return
}

// 获取单个币种
func GetCoinInfo(ctx context.Context, in *npool.GetCoinInfoRequest) (resp *npool.GetCoinInfoResponse, err error) {
	tmpID := uuid.MustParse(in.GetID())
	entResp, err := db.Client().CoinInfo.Query().
		Where(coininfo.ID(tmpID)).
		Only(ctx)
	if entResp == nil || err != nil {
		err = status.Errorf(codes.NotFound, "record not found, err: %v", err)
		resp = nil
		return
	}
	resp = &npool.GetCoinInfoResponse{
		Info: dbRowToCoinInfo(entResp),
	}
	return
}

// 注册币种
func RegisterCoin(ctx context.Context, in *npool.CreateCoinInfoRequest) (resp *npool.CreateCoinInfoResponse, err error) {
	resp = &npool.CreateCoinInfoResponse{}
	entResp, err := db.Client().CoinInfo.Query().
		Where(
			coininfo.Name(in.Info.Name),
		).Only(ctx)
	if err == nil {
		if in.Info.Unit == entResp.Unit {
			resp.Info = dbRowToCoinInfo(entResp)
			err = nil
		} else {
			err = status.Errorf(codes.AlreadyExists, "coin name already registered as: %v, unit: %v", entResp.Name, entResp.Unit)
		}
		return
	}
	fmt.Printf("inside crud: gonna create")
	// MARK 默认均为在售商品？
	tmpID := uuid.New()
	if len(in.Info.ID) > 0 {
		tmpID = uuid.MustParse(in.Info.ID)
	}
	entResp, err = db.Client().CoinInfo.Create().
		SetID(tmpID).
		SetCoinTypeID(in.Info.Enum).
		SetName(in.Info.Name).
		SetUnit(in.Info.Unit).
		SetIsPresale(false).
		Save(ctx)
	if err == nil {
		resp.Info = dbRowToCoinInfo(entResp)
	}
	return resp, err
}

// 设置币种权限
func SetCoinPresale(ctx context.Context, in *npool.UpdateCoinInfoRequest) (resp *npool.UpdateCoinInfoResponse, err error) {
	resp = &npool.UpdateCoinInfoResponse{}
	ci, err := db.Client().CoinInfo.Query().
		Where(
			coininfo.And(
				coininfo.ID(uuid.MustParse(in.Info.GetID())),
			),
		).First(ctx)
	if ci == nil || err != nil {
		err = status.Errorf(codes.NotFound, "no record found, err: %v", err)
		return
	}
	ci, err = ci.Update().
		SetIsPresale(in.Info.PreSale).
		Save(ctx)
	if err == nil {
		resp.Info = dbRowToCoinInfo(ci)
	}
	return
}
