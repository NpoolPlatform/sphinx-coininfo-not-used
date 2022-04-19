package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type CoinInfo struct {
	ent.Schema
}

func (CoinInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (CoinInfo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("name").NotEmpty().Unique(),
		field.String("unit").NotEmpty().Default(""),
		field.Uint64("reserved_amount").
			Default(0),
		field.Bool("pre_sale").Default(false),
		field.String("logo").Default(""),
		field.String("env").Default("main"), // main or test
		field.Bool("for_pay").Default(false),
		field.String("home_page").Default(""),
		field.String("specs").Default(""),
	}
}

func (CoinInfo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").
			Unique(),
	}
}
