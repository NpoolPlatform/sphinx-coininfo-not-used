// Code generated by entc, DO NOT EDIT.

package coininfo

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CoinTypeID applies equality check predicate on the "coin_type_id" field. It's identical to CoinTypeIDEQ.
func CoinTypeID(v int32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Unit applies equality check predicate on the "unit" field. It's identical to UnitEQ.
func Unit(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnit), v))
	})
}

// IsPresale applies equality check predicate on the "is_presale" field. It's identical to IsPresaleEQ.
func IsPresale(v bool) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsPresale), v))
	})
}

// LogoImage applies equality check predicate on the "logo_image" field. It's identical to LogoImageEQ.
func LogoImage(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLogoImage), v))
	})
}

// CoinTypeIDEQ applies the EQ predicate on the "coin_type_id" field.
func CoinTypeIDEQ(v int32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDNEQ applies the NEQ predicate on the "coin_type_id" field.
func CoinTypeIDNEQ(v int32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDIn applies the In predicate on the "coin_type_id" field.
func CoinTypeIDIn(vs ...int32) predicate.CoinInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDNotIn applies the NotIn predicate on the "coin_type_id" field.
func CoinTypeIDNotIn(vs ...int32) predicate.CoinInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDGT applies the GT predicate on the "coin_type_id" field.
func CoinTypeIDGT(v int32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDGTE applies the GTE predicate on the "coin_type_id" field.
func CoinTypeIDGTE(v int32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLT applies the LT predicate on the "coin_type_id" field.
func CoinTypeIDLT(v int32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLTE applies the LTE predicate on the "coin_type_id" field.
func CoinTypeIDLTE(v int32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoinTypeID), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.CoinInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.CoinInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// UnitEQ applies the EQ predicate on the "unit" field.
func UnitEQ(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnit), v))
	})
}

// UnitNEQ applies the NEQ predicate on the "unit" field.
func UnitNEQ(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUnit), v))
	})
}

// UnitIn applies the In predicate on the "unit" field.
func UnitIn(vs ...string) predicate.CoinInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUnit), v...))
	})
}

// UnitNotIn applies the NotIn predicate on the "unit" field.
func UnitNotIn(vs ...string) predicate.CoinInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUnit), v...))
	})
}

// UnitGT applies the GT predicate on the "unit" field.
func UnitGT(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUnit), v))
	})
}

// UnitGTE applies the GTE predicate on the "unit" field.
func UnitGTE(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUnit), v))
	})
}

// UnitLT applies the LT predicate on the "unit" field.
func UnitLT(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUnit), v))
	})
}

// UnitLTE applies the LTE predicate on the "unit" field.
func UnitLTE(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUnit), v))
	})
}

// UnitContains applies the Contains predicate on the "unit" field.
func UnitContains(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUnit), v))
	})
}

// UnitHasPrefix applies the HasPrefix predicate on the "unit" field.
func UnitHasPrefix(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUnit), v))
	})
}

// UnitHasSuffix applies the HasSuffix predicate on the "unit" field.
func UnitHasSuffix(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUnit), v))
	})
}

// UnitEqualFold applies the EqualFold predicate on the "unit" field.
func UnitEqualFold(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUnit), v))
	})
}

// UnitContainsFold applies the ContainsFold predicate on the "unit" field.
func UnitContainsFold(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUnit), v))
	})
}

// IsPresaleEQ applies the EQ predicate on the "is_presale" field.
func IsPresaleEQ(v bool) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsPresale), v))
	})
}

// IsPresaleNEQ applies the NEQ predicate on the "is_presale" field.
func IsPresaleNEQ(v bool) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsPresale), v))
	})
}

// LogoImageEQ applies the EQ predicate on the "logo_image" field.
func LogoImageEQ(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLogoImage), v))
	})
}

// LogoImageNEQ applies the NEQ predicate on the "logo_image" field.
func LogoImageNEQ(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLogoImage), v))
	})
}

// LogoImageIn applies the In predicate on the "logo_image" field.
func LogoImageIn(vs ...string) predicate.CoinInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLogoImage), v...))
	})
}

// LogoImageNotIn applies the NotIn predicate on the "logo_image" field.
func LogoImageNotIn(vs ...string) predicate.CoinInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLogoImage), v...))
	})
}

// LogoImageGT applies the GT predicate on the "logo_image" field.
func LogoImageGT(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLogoImage), v))
	})
}

// LogoImageGTE applies the GTE predicate on the "logo_image" field.
func LogoImageGTE(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLogoImage), v))
	})
}

// LogoImageLT applies the LT predicate on the "logo_image" field.
func LogoImageLT(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLogoImage), v))
	})
}

// LogoImageLTE applies the LTE predicate on the "logo_image" field.
func LogoImageLTE(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLogoImage), v))
	})
}

// LogoImageContains applies the Contains predicate on the "logo_image" field.
func LogoImageContains(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLogoImage), v))
	})
}

// LogoImageHasPrefix applies the HasPrefix predicate on the "logo_image" field.
func LogoImageHasPrefix(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLogoImage), v))
	})
}

// LogoImageHasSuffix applies the HasSuffix predicate on the "logo_image" field.
func LogoImageHasSuffix(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLogoImage), v))
	})
}

// LogoImageEqualFold applies the EqualFold predicate on the "logo_image" field.
func LogoImageEqualFold(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLogoImage), v))
	})
}

// LogoImageContainsFold applies the ContainsFold predicate on the "logo_image" field.
func LogoImageContainsFold(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLogoImage), v))
	})
}

// HasTransactions applies the HasEdge predicate on the "transactions" edge.
func HasTransactions() predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TransactionsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TransactionsTable, TransactionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTransactionsWith applies the HasEdge predicate on the "transactions" edge with a given conditions (other predicates).
func HasTransactionsWith(preds ...predicate.Transaction) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TransactionsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TransactionsTable, TransactionsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReviews applies the HasEdge predicate on the "reviews" edge.
func HasReviews() predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReviewsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ReviewsTable, ReviewsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReviewsWith applies the HasEdge predicate on the "reviews" edge with a given conditions (other predicates).
func HasReviewsWith(preds ...predicate.Review) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReviewsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ReviewsTable, ReviewsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWalletNodes applies the HasEdge predicate on the "wallet_nodes" edge.
func HasWalletNodes() predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WalletNodesTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, WalletNodesTable, WalletNodesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWalletNodesWith applies the HasEdge predicate on the "wallet_nodes" edge with a given conditions (other predicates).
func HasWalletNodesWith(preds ...predicate.WalletNode) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WalletNodesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, WalletNodesTable, WalletNodesPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CoinInfo) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CoinInfo) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CoinInfo) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		p(s.Not())
	})
}
