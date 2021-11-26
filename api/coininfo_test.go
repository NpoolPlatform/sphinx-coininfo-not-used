package api

import (
	"testing"

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
	assert.Nil(t, err)
	assert.Equal(t, 200, retCreateCoinInfo.StatusCode())

	retGetCoinInfoRequest, err := testaio.UnifyRestyQuery("/v1/get/coin/info", &npool.GetCoinInfoRequest{
		ID: testaio.CoinInfo.ID,
	})
	assert.Nil(t, err)
	assert.Equal(t, 200, retGetCoinInfoRequest.StatusCode())

	retGetCoinInfosRequest, err := testaio.UnifyRestyQuery("/v1/get/coin/infos", &emptypb.Empty{})
	assert.Nil(t, err)
	assert.Equal(t, 200, retGetCoinInfosRequest.StatusCode())

	retUpdateCoinInfo, err := testaio.UnifyRestyQuery("/v1/update/coin/info", &npool.UpdateCoinInfoRequest{
		Info: testaio.CoinInfo,
	})
	assert.Nil(t, err)
	assert.Equal(t, 200, retUpdateCoinInfo.StatusCode())
}
