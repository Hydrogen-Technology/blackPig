// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"service/internal/ent/times"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Times is the model entity for the Times schema.
type Times struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// UserId holds the value of the "userId" field.
	UserId int64 `json:"userId,omitempty"`
	// TimeSpace holds the value of the "timeSpace" field.
	TimeSpace string `json:"timeSpace,omitempty"`
	// StartTime holds the value of the "startTime" field.
	StartTime    string `json:"startTime,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Times) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case times.FieldID, times.FieldUserId:
			values[i] = new(sql.NullInt64)
		case times.FieldTimeSpace, times.FieldStartTime:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Times fields.
func (t *Times) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case times.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int64(value.Int64)
		case times.FieldUserId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field userId", values[i])
			} else if value.Valid {
				t.UserId = value.Int64
			}
		case times.FieldTimeSpace:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field timeSpace", values[i])
			} else if value.Valid {
				t.TimeSpace = value.String
			}
		case times.FieldStartTime:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field startTime", values[i])
			} else if value.Valid {
				t.StartTime = value.String
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Times.
// This includes values selected through modifiers, order, etc.
func (t *Times) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// Update returns a builder for updating this Times.
// Note that you need to call Times.Unwrap() before calling this method if this Times
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Times) Update() *TimesUpdateOne {
	return NewTimesClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Times entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Times) Unwrap() *Times {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Times is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Times) String() string {
	var builder strings.Builder
	builder.WriteString("Times(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("userId=")
	builder.WriteString(fmt.Sprintf("%v", t.UserId))
	builder.WriteString(", ")
	builder.WriteString("timeSpace=")
	builder.WriteString(t.TimeSpace)
	builder.WriteString(", ")
	builder.WriteString("startTime=")
	builder.WriteString(t.StartTime)
	builder.WriteByte(')')
	return builder.String()
}

// TimesSlice is a parsable slice of Times.
type TimesSlice []*Times
