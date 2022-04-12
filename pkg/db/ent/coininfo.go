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
	// ReservedAmount holds the value of the "reserved_amount" field.
	ReservedAmount uint64 `json:"reserved_amount,omitempty"`
	// PreSale holds the value of the "pre_sale" field.
	PreSale bool `json:"pre_sale,omitempty"`
	// Logo holds the value of the "logo" field.
	Logo string `json:"logo,omitempty"`
	// Env holds the value of the "env" field.
	Env string `json:"env,omitempty"`
	// ForPay holds the value of the "for_pay" field.
	ForPay bool `json:"for_pay,omitempty"`
	// HomePage holds the value of the "home_page" field.
	HomePage string `json:"home_page,omitempty"`
	// Specs holds the value of the "specs" field.
	Specs string `json:"specs,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CoinInfo) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case coininfo.FieldPreSale, coininfo.FieldForPay:
			values[i] = new(sql.NullBool)
		case coininfo.FieldReservedAmount, coininfo.FieldCreatedAt, coininfo.FieldUpdatedAt, coininfo.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case coininfo.FieldName, coininfo.FieldUnit, coininfo.FieldLogo, coininfo.FieldEnv, coininfo.FieldHomePage, coininfo.FieldSpecs:
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
		case coininfo.FieldReservedAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field reserved_amount", values[i])
			} else if value.Valid {
				ci.ReservedAmount = uint64(value.Int64)
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
		case coininfo.FieldEnv:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field env", values[i])
			} else if value.Valid {
				ci.Env = value.String
			}
		case coininfo.FieldForPay:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field for_pay", values[i])
			} else if value.Valid {
				ci.ForPay = value.Bool
			}
		case coininfo.FieldHomePage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field home_page", values[i])
			} else if value.Valid {
				ci.HomePage = value.String
			}
		case coininfo.FieldSpecs:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field specs", values[i])
			} else if value.Valid {
				ci.Specs = value.String
			}
		case coininfo.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ci.CreatedAt = uint32(value.Int64)
			}
		case coininfo.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ci.UpdatedAt = uint32(value.Int64)
			}
		case coininfo.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ci.DeletedAt = uint32(value.Int64)
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
	builder.WriteString(", reserved_amount=")
	builder.WriteString(fmt.Sprintf("%v", ci.ReservedAmount))
	builder.WriteString(", pre_sale=")
	builder.WriteString(fmt.Sprintf("%v", ci.PreSale))
	builder.WriteString(", logo=")
	builder.WriteString(ci.Logo)
	builder.WriteString(", env=")
	builder.WriteString(ci.Env)
	builder.WriteString(", for_pay=")
	builder.WriteString(fmt.Sprintf("%v", ci.ForPay))
	builder.WriteString(", home_page=")
	builder.WriteString(ci.HomePage)
	builder.WriteString(", specs=")
	builder.WriteString(ci.Specs)
	builder.WriteString(", created_at=")
	builder.WriteString(fmt.Sprintf("%v", ci.CreatedAt))
	builder.WriteString(", updated_at=")
	builder.WriteString(fmt.Sprintf("%v", ci.UpdatedAt))
	builder.WriteString(", deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", ci.DeletedAt))
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
