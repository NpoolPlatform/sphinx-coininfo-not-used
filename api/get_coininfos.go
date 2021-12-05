package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/crud/coininfo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TODO check pagesize max size
func (s *Server) GetCoinInfos(ctx context.Context, in *npool.GetCoinInfosRequest) (*npool.GetCoinInfosResponse, error) {
	resp, total, err := coininfo.GetAllCoinInfos(ctx, coininfo.GetAllCoinInfosParams{})
	if err != nil {
		logger.Sugar().Errorf("GetCoinInfos call GetAllCoinInfos error %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	infos := make([]*npool.CoinInfo, len(resp))
	for i, info := range resp {
		infos[i] = &npool.CoinInfo{
			ID:        info.ID.String(),
			PreSale:   info.PreSale,
			Name:      info.Name,
			Unit:      info.Unit,
			Logo:      info.Logo,
			CreatedAt: info.CreatedAt,
			UpdatedAt: info.UpdatedAt,
		}
	}

	return &npool.GetCoinInfosResponse{
		Total: int32(total),
		Infos: infos,
	}, nil
}
