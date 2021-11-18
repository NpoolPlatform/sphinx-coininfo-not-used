package api

import (
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"

	testinit "github.com/NpoolPlatform/sphinx-coininfo/pkg/test-init"
)

var (
	tmpCoinInfo     npool.CoinInfoRow
	testInitAlready bool
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	tmpCoinInfo.CoinTypeID = 0
	tmpCoinInfo.CoinType = 0
	tmpCoinInfo.IsPresale = false
	tmpCoinInfo.Name = "Unknown"
	tmpCoinInfo.Unit = "DK"
}

func runByGithub() bool {
	runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION"))
	if err == nil && runByGithubAction {
		return true
	}
	if testInitAlready == false {
		testInitAlready = true
		err = testinit.Init()
		if err != nil {
			logger.Sugar().Errorf("test init failed: %v", err)
		}
	}
	return err == nil
}

func TestGetCoinInfo(t *testing.T) {
	if runByGithub() {
		return
	}
	cli := resty.New()
	_, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetCoinInfoRequest{
			CoinType:   tmpCoinInfo.CoinType,
			CoinTypeID: tmpCoinInfo.CoinTypeID,
		}).
		Get("http://localhost:50130/v1/coin/single")
	assert.Nil(t, err)
}

func TestGetCoinInfos(t *testing.T) {
	if runByGithub() {
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
	if runByGithub() {
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
	if runByGithub() {
		return
	}
	cli := resty.New()
	_, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.SetCoinPresaleRequest{
			CoinType:   tmpCoinInfo.CoinType,
			CoinTypeID: tmpCoinInfo.CoinTypeID,
			IsPresale:  false,
		}).
		Post("http://localhost:50130/v1/coin/presale")
	assert.Nil(t, err)
}
