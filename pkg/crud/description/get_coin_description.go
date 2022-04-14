package description

import (
	"context"

	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/description"
	"github.com/google/uuid"
)

func GetCoinDescriptionByCoinID(ctx context.Context, cid uuid.UUID) (coinDescription *ent.Description, err error) {
	client, err := db.Client()
	if err != nil {
		return nil, err
	}
	coinDescription, err = client.Description.Query().
		Where(description.CoinTypeIDEQ(cid)).
		Only(ctx)
	return
}

func ExistCoinDescriptionByCoinID(ctx context.Context, cid uuid.UUID) (bool, error) {
	client, err := db.Client()
	if err != nil {
		return false, err
	}
	return client.Description.Query().
		Where(description.CoinTypeIDEQ(cid)).
		Exist(ctx)
}
