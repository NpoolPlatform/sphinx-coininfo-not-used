package crud

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/message/npool/coininfo" //nolint
	testinit "github.com/NpoolPlatform/sphinx-coininfo/pkg/test-init"
	"github.com/stretchr/testify/assert"
)

var (
	tmpCoinInfo     npool.CoinInfo
	testInitAlready bool
)

func initStruct() {
	tmpCoinInfo.Enum = 0
	tmpCoinInfo.PreSale = false
	tmpCoinInfo.Name = "Unknown"
	tmpCoinInfo.Unit = "DK"
	tmpCoinInfo.ID = "6ba7b812-9dad-11d1-80b4-00c04fd430c8"
}

func init() {
	if runByGithub() {
		return
	}
	_, err := RegisterCoin(context.Background(), &npool.CreateCoinInfoRequest{
		Info: &npool.CoinInfo{
			ID:   tmpCoinInfo.ID,
			Enum: tmpCoinInfo.Enum,
			Name: tmpCoinInfo.Name,
			Unit: tmpCoinInfo.Unit,
		},
	})
	if err != nil {
		panic("create test coin failed")
	}
}

func runByGithub() bool {
	initStruct()
	runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION"))
	if err == nil && runByGithubAction {
		return true
	}
	if testInitAlready == false {
		testInitAlready = true
		err = testinit.Init()
		if err != nil {
			fmt.Printf("test init failed: %v", err)
		}
	}
	return err == nil
}

func TestGetCoinInfos(t *testing.T) {
	if runByGithub() {
		return
	}
	resp, err := GetCoinInfos(context.Background(), &npool.GetCoinInfosRequest{})
	if err != nil {
		panic(err)
	}
	assert.NotNil(t, resp)
}

func TestRegisterCoin(t *testing.T) {
	if runByGithub() {
		return
	}
	resp, err := RegisterCoin(context.Background(), &npool.CreateCoinInfoRequest{
		Info: &npool.CoinInfo{
			ID:   tmpCoinInfo.ID,
			Enum: tmpCoinInfo.Enum,
			Name: tmpCoinInfo.Name,
			Unit: tmpCoinInfo.Unit,
		},
	})
	if err != nil {
		panic(err)
	}
	assert.NotNil(t, resp)
}

func TestGetCoinInfo(t *testing.T) {
	if runByGithub() {
		return
	}
	resp, err := GetCoinInfo(context.Background(), &npool.GetCoinInfoRequest{
		ID: tmpCoinInfo.ID,
	})
	if err != nil {
		assert.Nil(t, resp)
	} else {
		assert.NotNil(t, resp)
		assert.Equal(t, tmpCoinInfo.Name, resp.Info.Name)
	}
}

func TestSetCoinPresale(t *testing.T) {
	if runByGithub() {
		return
	}
	resp, err := RegisterCoin(context.Background(), &npool.CreateCoinInfoRequest{
		Info: &npool.CoinInfo{
			ID:   tmpCoinInfo.ID,
			Enum: tmpCoinInfo.Enum,
			Name: tmpCoinInfo.Name,
			Unit: tmpCoinInfo.Unit,
		},
	})
	if err != nil {
		panic(err)
	}
	assert.NotNil(t, resp)
	assert.Equal(t, tmpCoinInfo.Unit, resp.Info.Unit)
	respNew, err := SetCoinPresale(context.Background(), &npool.UpdateCoinInfoRequest{
		Info: &npool.CoinInfo{
			ID:      tmpCoinInfo.ID,
			Enum:    tmpCoinInfo.Enum,
			Name:    tmpCoinInfo.Name,
			Unit:    tmpCoinInfo.Unit,
			PreSale: false,
		},
	})
	if err != nil {
		panic(err)
	}
	assert.Equal(t, false, respNew.Info.PreSale)
}
