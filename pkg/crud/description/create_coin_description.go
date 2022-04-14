package description

import (
	"context"

	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	"github.com/google/uuid"
)

type CreateCoinDescriptionParam struct {
	CoinTypeID uuid.UUID
	Title      string
	Message    string
	UsedFor    string
}

func CreateCoinDescription(ctx context.Context, info *CreateCoinDescriptionParam) (*ent.Description, error) {
	client, err := db.Client()
	if err != nil {
		return nil, err
	}
	client.Description.Create()
	return client.Description.Create().
		SetCoinTypeID(info.CoinTypeID).
		SetTitle(info.Title).
		SetMessage(info.Message).
		SetUsedFor(info.UsedFor).
		Save(ctx)
}
