package api

import (
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"

	testinit "github.com/NpoolPlatform/sphinx-coininfo/pkg/test-init"
)

var tmpCoinInfo *npool.CoinInfoRow

func init() {
	err := testinit.Init()
	if err != nil {
		panic(err)
	}
	tmpCoinInfo = &npool.CoinInfoRow{
		CoinType:  -2,
		IsPresale: false,
		Name:      "Unit Test",
		Unit:      "UT",
	}
}

func TestGetCoinInfo(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	cli := resty.New()
	_, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetCoinInfoRequest{
			CoinType: tmpCoinInfo.CoinType,
		}).
		Get("http://localhost:50130/v1/coin/single")
	assert.Nil(t, err)
}

func TestGetCoinInfos(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	cli := resty.New()
	_, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetCoinInfosRequest{}).
		Get("http://localhost:50130/v1/coin/infos")
	assert.Nil(t, err)
}

func TestRegisterCoin(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	cli := resty.New()
	_, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.RegisterCoinRequest{
			CoinType: tmpCoinInfo.CoinType,
			Name:     tmpCoinInfo.Name,
			Unit:     tmpCoinInfo.Unit,
		}).
		Post("http://localhost:50130/v1/coin/register")
	assert.Nil(t, err)
}

func TestSetCoinPresale(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	cli := resty.New()
	_, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.SetCoinPresaleRequest{
			CoinType:  tmpCoinInfo.CoinType,
			IsPresale: false,
		}).
		Post("http://localhost:50130/v1/coin/presale")
	assert.Nil(t, err)
}
