package api

import (
	"encoding/json"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/coininfo" //nolint
	"github.com/stretchr/testify/assert"
	"golang.org/x/xerrors"
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
	var err error

	// 注册币种
	retCreateCoinInfo := &npool.CreateCoinInfoResponse{}
	err = testOneUnit("/v1/create/coin/info", &npool.CreateCoinInfoRequest{
		Info: testaio.CoinInfo,
	}, retCreateCoinInfo)
	assert.Nil(t, err)

	// 获取币种
	retGetCoinInfo := &npool.GetCoinInfoResponse{}
	err = testOneUnit("/v1/get/coin/info", &npool.GetCoinInfoRequest{
		ID: testaio.CoinInfo.ID,
	}, retGetCoinInfo)
	assert.Nil(t, err)

	// 获取币种列表
	retGetCoinInfosResp := &npool.GetCoinInfosResponse{}
	err = testOneUnit("/v1/get/coin/infos", &emptypb.Empty{}, retGetCoinInfosResp)
	assert.Nil(t, err)

	// 进行设置
	retUpdateCoinInfo := &npool.UpdateCoinInfoResponse{}
	err = testOneUnit("/v1/update/coin/info", &npool.UpdateCoinInfoRequest{
		Info: testaio.CoinInfo,
	}, retUpdateCoinInfo)
	assert.Nil(t, err)
}

func testOneUnit(path string, req interface{ String() string }, targetStruct interface{}) (err error) {
	restyResponse, err := testaio.UnifyRestyQuery(path, req)
	logger.Sugar().Infof(path+"%v", req.String())
	if err != nil {
		return
	}
	if restyResponse.StatusCode() != 200 {
		err = xerrors.Errorf(path+" failed, code %v", restyResponse.StatusCode())
		return
	}
	err = json.Unmarshal(restyResponse.Body(), targetStruct)
	if err != nil {
		err = xerrors.New(path + " response unmarshal failed! got: " + restyResponse.String())
	}
	return
}
