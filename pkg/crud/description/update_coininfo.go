package description

import (
	"context"

	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	"github.com/google/uuid"
)

func UpdateCoinDescriptionByID(ctx context.Context, id, coinID uuid.UUID, usedFor, title, message string) (*ent.Description, error) {
	client, err := db.Client()
	if err != nil {
		return nil, err
	}

	return client.
		Description.
		UpdateOneID(id).
		SetCoinTypeID(coinID).
		SetUsedFor(usedFor).
		SetTitle(title).
		SetMessage(message).
		Save(ctx)
}
