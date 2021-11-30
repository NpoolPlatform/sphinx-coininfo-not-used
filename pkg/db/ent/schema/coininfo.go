package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type CoinInfo struct {
	ent.Schema
}

func (CoinInfo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("name").NotEmpty().Unique(),
		field.String("unit").NotEmpty().Default(""),
		field.Bool("pre_sale").Default(false),
		field.String("logo").Default(""),
	}
}

func (CoinInfo) Indexs() []ent.Index {
	return []ent.Index{}
}
