package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/sphinx-coininfo/pkg/types"
)

type Description struct {
	ent.Schema
}

func (Description) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("coin_id", uuid.UUID{}).Unique(),
		field.String("human_readable_name"),
		field.JSON("descriptions", []types.CoinDescription{}),
		field.String("spec_title"),
		field.JSON("specs", []types.CoinSpec{}),
		field.Uint32("created_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("updated_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("deleted_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}
