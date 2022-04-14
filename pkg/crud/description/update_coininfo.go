package description

import (
	"context"

	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	"github.com/google/uuid"
)

func UpdateCoinDescriptionByID(ctx context.Context, title, message, usedFor string, id uuid.UUID) (*ent.Description, error) {
	client, err := db.Client()
	if err != nil {
		return nil, err
	}

	return client.
		Description.
		UpdateOneID(id).
		SetTitle(title).
		SetMessage(message).
		SetUsedFor(usedFor).
		Save(ctx)
}
