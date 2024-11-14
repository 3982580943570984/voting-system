package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Complaint holds the schema definition for the Complaint entity.
type Complaint struct {
	ent.Schema
}

// Fields of the Complaint.
func (Complaint) Fields() []ent.Field {
	return []ent.Field{
		field.String("content"),

		// TODO: consider using enum for status
		field.String("status").
			Default("pending"),

		field.Time("timestamp").
			Default(time.Now),
	}
}

// Edges of the Complaint.
func (Complaint) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("complaints").
			Unique().
			Required(),

		edge.From("election", Election.Type).
			Ref("complaints").
			Unique().
			Required(),

		edge.From("comment", Comment.Type).
			Ref("complaints").
			Unique().
			Required(),
	}
}
