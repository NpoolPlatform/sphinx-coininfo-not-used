package core

import (
	"context"

	"github.com/NpoolPlatform/sphinx-coininfo/message/npool"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	"google.golang.org/protobuf/types/known/emptypb"
)

var Client *ent.Client

func init() {
	Client = db.Client()
}

// 查询全部币种
func GetCoinInfos(ctx context.Context, _ *emptypb.Empty) (resp *npool.GetCoinInfosResponse, err error) {
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
