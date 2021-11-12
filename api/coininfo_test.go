package api

import (
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/sphinx-coininfo/message/npool"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetCoinInfos(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	cli := resty.New()
	_, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetCoinInfosRequest{}).
		Get("http://localhost:36759/v0/coin/infos")
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
			NeedSigninfo: true,
			Name:         "Filecoin",
			Unit:         "FIL",
		}).
		Post("http://localhost:36759/v0/coin/register")
	assert.Nil(t, err)
}
