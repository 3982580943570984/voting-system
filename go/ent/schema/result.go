package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Result holds the schema definition for the result entity.
type Result struct {
	ent.Schema
}

// Fields of the Result.
func (Result) Fields() []ent.Field {
	return []ent.Field{
		field.Int("votes").
			Default(0),
	}
}

// Edges of the Result.
func (Result) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("election", Election.Type).
			Ref("results").
			Unique().
			Required(),

		edge.From("candidate", Candidate.Type).
			Ref("results").
			Unique().
			Required(),
	}
}
