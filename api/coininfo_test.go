package api

import (
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/go-resty/resty/v2"

	testinit "github.com/NpoolPlatform/sphinx-coininfo/pkg/test-init"
)

var (
	tmpCoinInfo     npool.CoinInfo
	testInitAlready bool
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	tmpCoinInfo.Enum = 0
	tmpCoinInfo.ID = "6ba7b812-9dad-11d1-80b4-00c04fd430c8"
	tmpCoinInfo.PreSale = false
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
			ID: tmpCoinInfo.ID,
		}).
		Get("http://localhost:50150/v1/coin/single")
	if err != nil {
		panic(err)
	}
}

func TestGetCoinInfos(t *testing.T) {
	if runByGithub() {
		return
	}
	cli := resty.New()
	_, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetCoinInfosRequest{}).
		Get("http://localhost:50150/v1/coin/infos")
	if err != nil {
		panic(err)
	}
}

func TestRegisterCoin(t *testing.T) {
	if runByGithub() {
		return
	}
	cli := resty.New()
	_, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateCoinInfoRequest{
			Info: &npool.CoinInfo{
				ID:   tmpCoinInfo.ID,
				Enum: tmpCoinInfo.Enum,
				Name: tmpCoinInfo.Name,
				Unit: tmpCoinInfo.Unit,
			},
		}).
		Post("http://localhost:50150/v1/coin/register")
	if err != nil {
		panic(err)
	}
}

func TestSetCoinPresale(t *testing.T) {
	if runByGithub() {
		return
	}
	cli := resty.New()
	_, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UpdateCoinInfoRequest{
			Info: &npool.CoinInfo{
				Enum:    tmpCoinInfo.Enum,
				ID:      tmpCoinInfo.ID,
				PreSale: false,
			},
		}).
		Post("http://localhost:50150/v1/coin/presale")
	if err != nil {
		panic(err)
	}
}
