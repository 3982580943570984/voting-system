package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Feedback holds the schema definition for the Feedback entity.
type Feedback struct {
	ent.Schema
}

// Fields of the Feedback.
func (Feedback) Fields() []ent.Field {
	return []ent.Field{
		field.String("content"),

		field.Int("rating").
			Min(1).
			Max(5),

		field.Time("timestamp").
			Default(time.Now),
	}
}

// Edges of the Feedback.
func (Feedback) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("feedbacks").
			Unique().
			Required(),
	}
}
