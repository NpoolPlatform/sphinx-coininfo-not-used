// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/coininfo"
	"github.com/google/uuid"
)

// CoinInfo is the model entity for the CoinInfo schema.
type CoinInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Unit holds the value of the "unit" field.
	Unit string `json:"unit,omitempty"`
	// PreSale holds the value of the "pre_sale" field.
	PreSale bool `json:"pre_sale,omitempty"`
	// Logo holds the value of the "logo" field.
	Logo string `json:"logo,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CoinInfo) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case coininfo.FieldPreSale:
			values[i] = new(sql.NullBool)
		case coininfo.FieldName, coininfo.FieldUnit, coininfo.FieldLogo:
			values[i] = new(sql.NullString)
		case coininfo.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CoinInfo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CoinInfo fields.
func (ci *CoinInfo) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case coininfo.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ci.ID = *value
			}
		case coininfo.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ci.Name = value.String
			}
		case coininfo.FieldUnit:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field unit", values[i])
			} else if value.Valid {
				ci.Unit = value.String
			}
		case coininfo.FieldPreSale:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field pre_sale", values[i])
			} else if value.Valid {
				ci.PreSale = value.Bool
			}
		case coininfo.FieldLogo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field logo", values[i])
			} else if value.Valid {
				ci.Logo = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CoinInfo.
// Note that you need to call CoinInfo.Unwrap() before calling this method if this CoinInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (ci *CoinInfo) Update() *CoinInfoUpdateOne {
	return (&CoinInfoClient{config: ci.config}).UpdateOne(ci)
}

// Unwrap unwraps the CoinInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ci *CoinInfo) Unwrap() *CoinInfo {
	tx, ok := ci.config.driver.(*txDriver)
	if !ok {
		panic("ent: CoinInfo is not a transactional entity")
	}
	ci.config.driver = tx.drv
	return ci
}

// String implements the fmt.Stringer.
func (ci *CoinInfo) String() string {
	var builder strings.Builder
	builder.WriteString("CoinInfo(")
	builder.WriteString(fmt.Sprintf("id=%v", ci.ID))
	builder.WriteString(", name=")
	builder.WriteString(ci.Name)
	builder.WriteString(", unit=")
	builder.WriteString(ci.Unit)
	builder.WriteString(", pre_sale=")
	builder.WriteString(fmt.Sprintf("%v", ci.PreSale))
	builder.WriteString(", logo=")
	builder.WriteString(ci.Logo)
	builder.WriteByte(')')
	return builder.String()
}

// CoinInfos is a parsable slice of CoinInfo.
type CoinInfos []*CoinInfo

func (ci CoinInfos) config(cfg config) {
	for _i := range ci {
		ci[_i].config = cfg
	}
}
