package crud

import (
	"context"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/coininfo"
	testinit "github.com/NpoolPlatform/sphinx-coininfo/pkg/test-init"
	"github.com/stretchr/testify/assert"
)

var (
	ctx         context.Context
	tmpCoinInfo npool.CoinInfoRow
	Flag删库      bool
)

func init() {
	if testinit.Init() != nil {
		panic("testinit failed")
	}
	Flag删库 = true
	ctx = context.Background()
	if Flag删库 {
		// dangerous
		_, err := db.Client().CoinInfo.Delete().Where(coininfo.Not(coininfo.Name("anything"))).Exec(ctx)
		if err != nil {
			panic(err)
		}
	}
	tmpCoinInfo.CoinType = 0
	tmpCoinInfo.IsPresale = false
	tmpCoinInfo.Name = "Unknown"
	tmpCoinInfo.Unit = "DK"
	RegisterCoin(ctx, &npool.RegisterCoinRequest{
		CoinType: tmpCoinInfo.CoinType,
		Name:     tmpCoinInfo.Name,
		Unit:     tmpCoinInfo.Unit,
	})
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
	if resp != nil {
		assert.Equal(t, tmpCoinInfo.Unit, resp.Unit)
		assert.Equal(t, tmpCoinInfo.CoinType, resp.CoinType)
	} else {
		assert.NotNil(t, err)
	}
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
