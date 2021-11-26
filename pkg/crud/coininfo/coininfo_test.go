package coininfo

import (
	"context"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	testinit "github.com/NpoolPlatform/sphinx-coininfo/pkg/test-init"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/testaio"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var entCoinInfo = &ent.CoinInfo{
	ID:         uuid.New(),
	CoinTypeID: 1,
	Name:       "FIL",
	Unit:       "FIL",
	IsPresale:  false,
	LogoImage:  "",
}

func TestCRUD(t *testing.T) {
	if testaio.RunByGithub() {
		return
	}
	if !testaio.InitAlready {
		assert.Nil(t, testinit.Init())
		testaio.InitAlready = true
	}
	ctx := context.Background()

	coinInfo, err := UpsertCoinInfoByName(ctx, entCoinInfo, entCoinInfo.Name)
	assert.Nil(t, err)
	logger.Sugar().Infof("CreateCoin result: %v", coinInfo)
	assert.True(t, assertCoinInfoEqual(entCoinInfo, coinInfo))

	coinInfo, err = GetCoinInfoByID(ctx, coinInfo.ID.String())
	assert.Nil(t, err)
	logger.Sugar().Infof("GetInfo result: %v", coinInfo)
	assert.True(t, assertCoinInfoEqual(entCoinInfo, coinInfo))

	coinInfos, err := GetAllCoinInfos(ctx)
	assert.Nil(t, err)
	logger.Sugar().Infof("GetInfos result: %v", coinInfos)
	assert.NotZero(t, len(coinInfos))

	coinInfo, err = UpdatePreSaleByID(ctx, false, coinInfo.ID.String())
	assert.Nil(t, err)
	logger.Sugar().Infof("UpdatePreSale result: %v", coinInfo)
	assert.False(t, coinInfo.IsPresale)
}

func assertCoinInfoEqual(coinInfoA, coinInfoB *ent.CoinInfo) bool {
	if coinInfoA == nil || coinInfoB == nil {
		return false
	}
	return coinInfoA.Name == coinInfoB.Name &&
		coinInfoA.Unit == coinInfoB.Unit &&
		coinInfoA.LogoImage == coinInfoB.LogoImage &&
		coinInfoA.IsPresale == coinInfoB.IsPresale
}
