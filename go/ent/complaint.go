// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"voting-system/ent/complaint"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Complaint is the model entity for the Complaint schema.
type Complaint struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Complaint) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case complaint.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Complaint fields.
func (c *Complaint) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case complaint.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Complaint.
// This includes values selected through modifiers, order, etc.
func (c *Complaint) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// Update returns a builder for updating this Complaint.
// Note that you need to call Complaint.Unwrap() before calling this method if this Complaint
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Complaint) Update() *ComplaintUpdateOne {
	return NewComplaintClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Complaint entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Complaint) Unwrap() *Complaint {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Complaint is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Complaint) String() string {
	var builder strings.Builder
	builder.WriteString("Complaint(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Complaints is a parsable slice of Complaint.
type Complaints []*Complaint
