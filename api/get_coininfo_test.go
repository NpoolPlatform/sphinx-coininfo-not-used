package api

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gopkg.in/resty.v1"
)

func TestGetCoinInfo(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	var (
		cli      = resty.New()
		coinInfo coininfo.CreateCoinInfoResponse
	)

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(coininfo.GetCoinInfoRequest{}).
		Post("http://localhost:36759/v1/create/coininfo")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &coinInfo)
		if assert.Nil(t, err) {
			assert.NotEqual(t, coinInfo.GetID(), uuid.UUID{}.String())
		}
	}
}
