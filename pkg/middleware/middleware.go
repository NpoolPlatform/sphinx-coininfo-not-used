package middleware

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/coininfo" //nolint
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/crud/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	"github.com/google/uuid"
)

func GetCoinInfo(ctx context.Context, in *npool.GetCoinInfoRequest) (resp *npool.GetCoinInfoResponse, err error) {
	coinInfo, err := coininfo.GetCoinInfoByID(ctx, in.ID)
	if err == nil {
		resp = &npool.GetCoinInfoResponse{
			Info: dbRowToCoinInfo(coinInfo),
		}
	}
	return
}

func GetCoinInfos(ctx context.Context, _ *npool.GetCoinInfosRequest) (resp *npool.GetCoinInfosResponse, err error) {
	entResp, err := coininfo.GetAllCoinInfos(ctx)
	var coinInfos []*npool.CoinInfo
	if err == nil {
		coinInfos = dbRowsToCoinInfos(entResp)
	}
	resp = &npool.GetCoinInfosResponse{
		Infos: coinInfos,
	}
	return
}

func CreateCoinInfo(ctx context.Context, in *npool.CreateCoinInfoRequest) (resp *npool.CreateCoinInfoResponse, err error) {
	coinInfo := &ent.CoinInfo{
		ID:         uuid.MustParse(in.Info.ID),
		CoinTypeID: in.Info.Enum,
		IsPresale:  in.Info.PreSale,
		Name:       in.Info.Name,
		Unit:       in.Info.Unit,
		LogoImage:  in.Info.LogoImage,
	}
	entResp, err := coininfo.UpsertCoinInfoByName(ctx, coinInfo, in.Info.Name)
	if err == nil {
		resp = &npool.CreateCoinInfoResponse{
			Info: dbRowToCoinInfo(entResp),
		}
	}
	return resp, err
}

func UpdateCoinInfo(ctx context.Context, in *npool.UpdateCoinInfoRequest) (resp *npool.UpdateCoinInfoResponse, err error) {
	coinInfo, err := coininfo.UpdatePreSaleByID(ctx, in.Info.PreSale, in.Info.ID)
	if err == nil {
		resp = &npool.UpdateCoinInfoResponse{
			Info: dbRowToCoinInfo(coinInfo),
		}
	}
	return
}

func dbRowToCoinInfo(row *ent.CoinInfo) *npool.CoinInfo {
	if row != nil {
		return &npool.CoinInfo{
			ID:        row.ID.String(),
			Enum:      row.CoinTypeID,
			Name:      row.Name,
			Unit:      row.Unit,
			PreSale:   row.IsPresale,
			LogoImage: row.LogoImage,
		}
	}
	return &npool.CoinInfo{}
}

func dbRowsToCoinInfos(entResp []*ent.CoinInfo) (coinInfos []*npool.CoinInfo) {
	coinInfos = make([]*npool.CoinInfo, len(entResp))
	for i, row := range entResp {
		coinInfos[i] = dbRowToCoinInfo(row)
	}
	return
}
