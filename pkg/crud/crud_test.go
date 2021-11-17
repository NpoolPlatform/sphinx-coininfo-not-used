package crud

import (
	"context"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/message/npool/coininfo"
	testinit "github.com/NpoolPlatform/sphinx-coininfo/pkg/test-init"
	"github.com/stretchr/testify/assert"
)

var (
	ctx         context.Context
	tmpCoinInfo *npool.CoinInfoRow
)

func init() {
	if testinit.Init() != nil {
		panic("testinit failed")
	}
	ctx = context.Background()
	tmpCoinInfo = &npool.CoinInfoRow{
		CoinType:  0,
		IsPresale: false,
		Name:      "Unit Test",
		Unit:      "UT",
	}
	TestRegisterCoin(&testing.T{})
}

func runByGithub() bool {
	runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION"))
	return err == nil && runByGithubAction
}

func TestGetCoinInfos(t *testing.T) {
	if runByGithub() {
		return
	}
	resp, err := GetCoinInfos(ctx, &npool.GetCoinInfosRequest{})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestRegisterCoin(t *testing.T) {
	if runByGithub() {
		return
	}
	resp, err := RegisterCoin(ctx, &npool.RegisterCoinRequest{
		CoinType: tmpCoinInfo.CoinType,
		Name:     tmpCoinInfo.Name,
		Unit:     tmpCoinInfo.Unit,
	})
	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.Equal(t, tmpCoinInfo.CoinType, resp.CoinType)
	assert.Equal(t, tmpCoinInfo.Unit, resp.Unit)
}

func TestGetCoinInfo(t *testing.T) {
	if runByGithub() {
		return
	}
	resp, err := GetCoinInfo(ctx, &npool.GetCoinInfoRequest{
		CoinType: tmpCoinInfo.CoinType,
	})
	if err != nil {
		assert.Nil(t, resp)
	} else {
		assert.NotNil(t, resp)
		assert.Equal(t, tmpCoinInfo.Name, resp.Name)
	}
}

func TestSetCoinPresale(t *testing.T) {
	if runByGithub() {
		return
	}
	resp, err := SetCoinPresale(ctx, &npool.SetCoinPresaleRequest{
		CoinType:  tmpCoinInfo.CoinType,
		IsPresale: !tmpCoinInfo.IsPresale,
	})
	assert.Nil(t, err)
	assert.Equal(t, tmpCoinInfo.IsPresale, !resp.IsPresale)
}
