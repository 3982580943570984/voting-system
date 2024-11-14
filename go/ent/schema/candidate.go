// schema/candidate.go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Candidate holds the schema definition for the Candidate entity.
type Candidate struct {
	ent.Schema
}

// Fields of the Candidate.
func (Candidate) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description").
			Optional(),
	}
}

// Edges of the Candidate.
func (Candidate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("election", Election.Type).
			Ref("candidates").
			Unique().
			Required(),
		edge.To("votes", Vote.Type),
		edge.To("results", ElectionResult.Type),
	}
}
