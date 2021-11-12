package core

import (
	"context"

	"github.com/NpoolPlatform/sphinx-coininfo/message/npool"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/coininfo"
)

var Client *ent.Client

func init() {
	Client = db.Client()
}

// 查询全部币种
func GetCoinInfos(ctx context.Context, _ *npool.GetCoinInfosRequest) (resp *npool.GetCoinInfosResponse, err error) {
	entResp, err := Client.CoinInfo.Query().All(ctx)
	tmpCIR := make([]*npool.CoinInfoRow, len(entResp))
	for i, row := range entResp {
		tmpCIR[i] = &npool.CoinInfoRow{
			Id:           row.ID,
			Name:         row.Name,
			Unit:         row.Unit,
			NeedSigninfo: row.NeedSigninfo,
		}
	}
	resp = &npool.GetCoinInfosResponse{
		Infos: tmpCIR,
	}
	return
}

// 注册币种
func RegisterCoin(ctx context.Context, in *npool.RegisterCoinRequest) (resp *npool.RegisterCoinResponse, err error) {
	entResp, err := Client.CoinInfo.Query().
		Where(
			coininfo.Name(in.Unit),
		).First(ctx)
	if entResp != nil {
		if in.Unit == entResp.Unit {
			resp = &npool.RegisterCoinResponse{Info: "success"}
			err = nil
		}
		return
	}
	// do create
	_, err = Client.CoinInfo.Create().
		SetName(in.Name).
		SetUnit(in.Unit).
		SetNeedSigninfo(in.NeedSigninfo).
		Save(ctx)
	if err == nil {
		resp = &npool.RegisterCoinResponse{Info: "success"}
	}
	return
}
