package api

import (
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/coininfo" //nolint
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"

	testinit "github.com/NpoolPlatform/sphinx-coininfo/pkg/test-init"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/testaio"
)

func TestCRUD(t *testing.T) {
	if testaio.RunByGithub() {
		return
	}
	if !testaio.InitAlready {
		assert.Nil(t, testinit.Init())
	}
	testaio.Host = "http://localhost:50150"

	retCreateCoinInfo, err := testaio.UnifyRestyQuery("/v1/create/coin/info", &npool.CreateCoinInfoRequest{
		Info: testaio.CoinInfo,
	})
	logger.Sugar().Infof("retCreateCoinInfo: %v", retCreateCoinInfo.String())
	assert.Equal(t, 200, retCreateCoinInfo.StatusCode())
	assert.Nil(t, err)

	retGetCoinInfoRequest, err := testaio.UnifyRestyQuery("/v1/get/coin/info", &npool.GetCoinInfoRequest{
		ID: testaio.CoinInfo.ID,
	})
	logger.Sugar().Infof("retGetCoinInfoRequest: %v", retGetCoinInfoRequest.String())
	assert.Equal(t, 200, retGetCoinInfoRequest.StatusCode())
	assert.Nil(t, err)

	retGetCoinInfosRequest, err := testaio.UnifyRestyQuery("/v1/get/coin/infos", &emptypb.Empty{})
	logger.Sugar().Infof("retGetCoinInfosRequest: %v", retGetCoinInfosRequest.String())
	assert.Equal(t, 200, retGetCoinInfosRequest.StatusCode())
	assert.Nil(t, err)

	retUpdateCoinInfo, err := testaio.UnifyRestyQuery("/v1/update/coin/info", &npool.UpdateCoinInfoRequest{
		Info: testaio.CoinInfo,
	})
	logger.Sugar().Infof("retUpdateCoinInfo: %v", retUpdateCoinInfo.String())
	assert.Equal(t, 200, retUpdateCoinInfo.StatusCode())
	assert.Nil(t, err)
}
