package app

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/coininfo" //nolint
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/crud"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	"github.com/google/uuid"
)

func GetCoinInfos(ctx context.Context, _ *npool.GetCoinInfosRequest) (resp *npool.GetCoinInfosResponse, err error) {
	entResp, err := crud.GetInfos(ctx)
	var coinInfos []*npool.CoinInfo
	if err == nil {
		coinInfos = dbRowsToCoinInfos(entResp)
	}
	resp = &npool.GetCoinInfosResponse{
		Infos: coinInfos,
	}
	return
}

func GetCoinInfo(ctx context.Context, in *npool.GetCoinInfoRequest) (resp *npool.GetCoinInfoResponse, err error) {
	coinInfo, err := crud.GetInfo(ctx, in.ID)
	if err == nil {
		resp = &npool.GetCoinInfoResponse{
			Info: dbRowToCoinInfo(coinInfo),
		}
	}
	return
}

func CreateCoinInfo(ctx context.Context, in *npool.CreateCoinInfoRequest) (resp *npool.CreateCoinInfoResponse, err error) {
	coinInfo := protoCoinInfoToDBRow(in.Info)
	coinInfo, err = crud.CreateCoin(ctx, coinInfo)
	if err == nil {
		resp = &npool.CreateCoinInfoResponse{
			Info: dbRowToCoinInfo(coinInfo),
		}
	}
	return resp, err
}

func UpdateCoinInfo(ctx context.Context, in *npool.UpdateCoinInfoRequest) (resp *npool.UpdateCoinInfoResponse, err error) {
	coinInfo := protoCoinInfoToDBRow(in.Info)
	coinInfo, err = crud.UpdatePreSale(ctx, coinInfo)
	if err == nil {
		resp = &npool.UpdateCoinInfoResponse{
			Info: dbRowToCoinInfo(coinInfo),
		}
	}
	return
}

func dbRowToCoinInfo(row *ent.CoinInfo) *npool.CoinInfo {
	return &npool.CoinInfo{
		ID:        row.ID.String(),
		Enum:      row.CoinTypeID,
		Name:      row.Name,
		Unit:      row.Unit,
		PreSale:   row.IsPresale,
		LogoImage: row.LogoImage,
	}
}

func dbRowsToCoinInfos(entResp []*ent.CoinInfo) (coinInfos []*npool.CoinInfo) {
	coinInfos = make([]*npool.CoinInfo, len(entResp))
	for i, row := range entResp {
		coinInfos[i] = dbRowToCoinInfo(row)
	}
	return
}

func protoCoinInfoToDBRow(in *npool.CoinInfo) (coinInfo *ent.CoinInfo) {
	var err error
	if in.Name != "" {
		coinInfo, err = crud.GetInfoByName(context.Background(), in.Name)
	} else {
		coinInfo, err = crud.GetInfo(context.Background(), in.ID)
	}
	if err == nil {
		coinInfo = &ent.CoinInfo{
			ID:         uuid.MustParse(in.ID),
			CoinTypeID: in.Enum,
			Name:       in.Name,
			Unit:       in.Unit,
			IsPresale:  in.PreSale,
			LogoImage:  in.LogoImage,
		}
	}
	return
}
