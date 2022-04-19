package description

import (
	"context"

	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/description"
	constant "github.com/NpoolPlatform/sphinx-coininfo/pkg/message/const"
	"github.com/google/uuid"
)

func ExistCoinDescriptionByID(ctx context.Context, id uuid.UUID) (bool, error) {
	client, err := db.Client()
	if err != nil {
		return false, err
	}

	return client.Description.
		Query().
		Where(description.ID(id)).
		Exist(ctx)
}

func GetCoinDescriptionByCoinID(ctx context.Context, cid uuid.UUID, limit, offset int32) (coinDescriptions []*ent.Description, total int, err error) {
	if limit == 0 {
		limit = constant.DefaultPageSize
	}

	client, err := db.Client()
	if err != nil {
		return nil, 0, err
	}

	total, err = client.Description.Query().Where(description.CoinTypeID(cid)).Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	coinDescriptions, err = client.Description.
		Query().
		Where(description.CoinTypeIDEQ(cid)).
		Order(ent.Desc(description.FieldUpdatedAt)).
		Limit(int(limit)).
		Offset(int(offset)).
		All(ctx)

	return coinDescriptions, total, err
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

func ExistCoinDescriptionByCoinIDAndUsedFor(ctx context.Context, cid uuid.UUID, usedFor string) (bool, error) {
	client, err := db.Client()
	if err != nil {
		return false, err
	}
	return client.Description.Query().
		Where(
			description.CoinTypeIDEQ(cid),
			description.UsedForEQ(usedFor),
		).
		Exist(ctx)
}
