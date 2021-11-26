package crud

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

func TestCRUD(t *testing.T) {
	if testaio.RunByGithub() {
		return
	}
	if !testaio.InitAlready {
		assert.Nil(t, testinit.Init())
		testaio.InitAlready = true
	}
	ctx := context.Background()
	coinInfo, err := CreateCoin(ctx, entCoinInfo)
	logger.Sugar().Infof("CreateCoin result: %v", coinInfo)
	AbortWhenError(t, err)
	coinInfo, err = GetInfo(ctx, coinInfo.ID.String())
	logger.Sugar().Infof("GetInfo result: %v", coinInfo)
	AbortWhenError(t, err)
	coinInfos, err := GetInfos(ctx)
	logger.Sugar().Infof("GetInfos result: %v", coinInfos)
	AbortWhenError(t, err)
	coinInfo.IsPresale = !coinInfo.IsPresale
	coinInfo, err = UpdatePreSale(ctx, coinInfo)
	logger.Sugar().Infof("UpdatePreSale result: %v", coinInfo)
	AbortWhenError(t, err)
}

func LogWhenError(err error) {
	if err != nil {
		logger.Sugar().Warn(err)
	}
}

func AbortWhenError(t *testing.T, err error) {
	if err != nil {
		logger.Sugar().Warn(err)
	}
	assert.Nil(t, err)
}

var entCoinInfo = &ent.CoinInfo{
	ID:         uuid.MustParse("8fbcbdc2-25ea-4ff0-b049-9d2f4c8ab646"),
	CoinTypeID: 1,
	Name:       "FIL",
	Unit:       "FIL",
	IsPresale:  false,
	LogoImage:  "",
}
