package client

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/coininfo"
	constant "github.com/NpoolPlatform/sphinx-coininfo/pkg/message/const"
)

func do(ctx context.Context, fn func(_ctx context.Context, cli npool.SphinxCoinInfoClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get coininfo connection: %v", err)
	}
	defer conn.Close()

	cli := npool.NewSphinxCoinInfoClient(conn)

	return fn(_ctx, cli)
}

func GetCoinInfos(ctx context.Context, conds cruder.FilterConds) ([]*npool.CoinInfo, error) {
	// conds: NOT USED NOW, will be used after refactor code
	infos, err := do(ctx, func(_ctx context.Context, cli npool.SphinxCoinInfoClient) (cruder.Any, error) {
		resp, err := cli.GetCoinInfos(ctx, &npool.GetCoinInfosRequest{
			Offset: 0,
			Limit:  100,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get coininfos: %v", err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get coininfos: %v", err)
	}
	return infos.([]*npool.CoinInfo), nil
}

func GetCoinInfo(ctx context.Context, id string) (*npool.CoinInfo, error) {
	// conds: NOT USED NOW, will be used after refactor code
	info, err := do(ctx, func(_ctx context.Context, cli npool.SphinxCoinInfoClient) (cruder.Any, error) {
		resp, err := cli.GetCoinInfo(ctx, &npool.GetCoinInfoRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get coininfo: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get coininfo: %v", err)
	}
	return info.(*npool.CoinInfo), nil
}
