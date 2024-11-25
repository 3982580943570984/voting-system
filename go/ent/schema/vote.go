package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Vote holds the schema definition for the Vote entity.
type Vote struct {
	ent.Schema
}

// Fields of the Vote.
func (Vote) Fields() []ent.Field {
	return []ent.Field{
		field.Time("timestamp").
			Default(time.Now),

		field.Bool("is_active").
			Default(true),
	}
}

// Edges of the Vote.
func (Vote) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("candidate", Candidate.Type).
			Ref("votes").
			Unique(),

		edge.From("user", User.Type).
			Ref("votes").
			Unique(),
	}
}
