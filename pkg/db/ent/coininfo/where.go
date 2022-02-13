// Code generated by entc, DO NOT EDIT.

package coininfo

import (
	"entgo.io/ent/dialect/sql"
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

// ReservedAmount applies equality check predicate on the "reserved_amount" field. It's identical to ReservedAmountEQ.
func ReservedAmount(v uint64) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReservedAmount), v))
	})
}

// PreSale applies equality check predicate on the "pre_sale" field. It's identical to PreSaleEQ.
func PreSale(v bool) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPreSale), v))
	})
}

// Logo applies equality check predicate on the "logo" field. It's identical to LogoEQ.
func Logo(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLogo), v))
	})
}

// Env applies equality check predicate on the "env" field. It's identical to EnvEQ.
func Env(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnv), v))
	})
}

// ForPay applies equality check predicate on the "for_pay" field. It's identical to ForPayEQ.
func ForPay(v bool) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldForPay), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
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

// ReservedAmountEQ applies the EQ predicate on the "reserved_amount" field.
func ReservedAmountEQ(v uint64) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReservedAmount), v))
	})
}

// ReservedAmountNEQ applies the NEQ predicate on the "reserved_amount" field.
func ReservedAmountNEQ(v uint64) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReservedAmount), v))
	})
}

// ReservedAmountIn applies the In predicate on the "reserved_amount" field.
func ReservedAmountIn(vs ...uint64) predicate.CoinInfo {
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
		s.Where(sql.In(s.C(FieldReservedAmount), v...))
	})
}

// ReservedAmountNotIn applies the NotIn predicate on the "reserved_amount" field.
func ReservedAmountNotIn(vs ...uint64) predicate.CoinInfo {
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
		s.Where(sql.NotIn(s.C(FieldReservedAmount), v...))
	})
}

// ReservedAmountGT applies the GT predicate on the "reserved_amount" field.
func ReservedAmountGT(v uint64) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReservedAmount), v))
	})
}

// ReservedAmountGTE applies the GTE predicate on the "reserved_amount" field.
func ReservedAmountGTE(v uint64) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReservedAmount), v))
	})
}

// ReservedAmountLT applies the LT predicate on the "reserved_amount" field.
func ReservedAmountLT(v uint64) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReservedAmount), v))
	})
}

// ReservedAmountLTE applies the LTE predicate on the "reserved_amount" field.
func ReservedAmountLTE(v uint64) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReservedAmount), v))
	})
}

// PreSaleEQ applies the EQ predicate on the "pre_sale" field.
func PreSaleEQ(v bool) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPreSale), v))
	})
}

// PreSaleNEQ applies the NEQ predicate on the "pre_sale" field.
func PreSaleNEQ(v bool) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPreSale), v))
	})
}

// LogoEQ applies the EQ predicate on the "logo" field.
func LogoEQ(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLogo), v))
	})
}

// LogoNEQ applies the NEQ predicate on the "logo" field.
func LogoNEQ(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLogo), v))
	})
}

// LogoIn applies the In predicate on the "logo" field.
func LogoIn(vs ...string) predicate.CoinInfo {
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
		s.Where(sql.In(s.C(FieldLogo), v...))
	})
}

// LogoNotIn applies the NotIn predicate on the "logo" field.
func LogoNotIn(vs ...string) predicate.CoinInfo {
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
		s.Where(sql.NotIn(s.C(FieldLogo), v...))
	})
}

// LogoGT applies the GT predicate on the "logo" field.
func LogoGT(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLogo), v))
	})
}

// LogoGTE applies the GTE predicate on the "logo" field.
func LogoGTE(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLogo), v))
	})
}

// LogoLT applies the LT predicate on the "logo" field.
func LogoLT(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLogo), v))
	})
}

// LogoLTE applies the LTE predicate on the "logo" field.
func LogoLTE(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLogo), v))
	})
}

// LogoContains applies the Contains predicate on the "logo" field.
func LogoContains(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLogo), v))
	})
}

// LogoHasPrefix applies the HasPrefix predicate on the "logo" field.
func LogoHasPrefix(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLogo), v))
	})
}

// LogoHasSuffix applies the HasSuffix predicate on the "logo" field.
func LogoHasSuffix(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLogo), v))
	})
}

// LogoEqualFold applies the EqualFold predicate on the "logo" field.
func LogoEqualFold(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLogo), v))
	})
}

// LogoContainsFold applies the ContainsFold predicate on the "logo" field.
func LogoContainsFold(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLogo), v))
	})
}

// EnvEQ applies the EQ predicate on the "env" field.
func EnvEQ(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnv), v))
	})
}

// EnvNEQ applies the NEQ predicate on the "env" field.
func EnvNEQ(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEnv), v))
	})
}

// EnvIn applies the In predicate on the "env" field.
func EnvIn(vs ...string) predicate.CoinInfo {
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
		s.Where(sql.In(s.C(FieldEnv), v...))
	})
}

// EnvNotIn applies the NotIn predicate on the "env" field.
func EnvNotIn(vs ...string) predicate.CoinInfo {
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
		s.Where(sql.NotIn(s.C(FieldEnv), v...))
	})
}

// EnvGT applies the GT predicate on the "env" field.
func EnvGT(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEnv), v))
	})
}

// EnvGTE applies the GTE predicate on the "env" field.
func EnvGTE(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEnv), v))
	})
}

// EnvLT applies the LT predicate on the "env" field.
func EnvLT(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEnv), v))
	})
}

// EnvLTE applies the LTE predicate on the "env" field.
func EnvLTE(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEnv), v))
	})
}

// EnvContains applies the Contains predicate on the "env" field.
func EnvContains(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEnv), v))
	})
}

// EnvHasPrefix applies the HasPrefix predicate on the "env" field.
func EnvHasPrefix(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEnv), v))
	})
}

// EnvHasSuffix applies the HasSuffix predicate on the "env" field.
func EnvHasSuffix(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEnv), v))
	})
}

// EnvEqualFold applies the EqualFold predicate on the "env" field.
func EnvEqualFold(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEnv), v))
	})
}

// EnvContainsFold applies the ContainsFold predicate on the "env" field.
func EnvContainsFold(v string) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEnv), v))
	})
}

// ForPayEQ applies the EQ predicate on the "for_pay" field.
func ForPayEQ(v bool) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldForPay), v))
	})
}

// ForPayNEQ applies the NEQ predicate on the "for_pay" field.
func ForPayNEQ(v bool) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldForPay), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.CoinInfo {
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
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.CoinInfo {
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
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.CoinInfo {
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
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.CoinInfo {
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
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.CoinInfo {
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
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.CoinInfo {
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
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.CoinInfo {
	return predicate.CoinInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
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
